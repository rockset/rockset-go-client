package rockset

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
)

// https://docs.rockset.com/rest-api/#documents

// AddDocuments adds new documents to a collection
//
// REST API documentation https://docs.rockset.com/rest-api/#adddocuments
func (rc *RockClient) AddDocuments(ctx context.Context, workspace, collection string,
	docs []interface{}) ([]openapi.DocumentStatus, error) {
	resp, err := rc.AddDocumentsWithOffset(ctx, workspace, collection, docs)
	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

// AddDocumentsWithOffset adds new documents to a collection, and returns the response with the offset(s),
// which can be used to wait until the collection includes the documents at or after the specified offset(s).
//
//	response, err := rs.AddDocumentsWithOffset(ctx, "commons", "users", docs)
//	if err != nil {
//	    return err
//	}
//	w := wait.New(rs)
//	err = w.UntilQueryable(ctx, "commons", "users", []string{response.GetLastOffset()})
//
// The added documents are now queryable in the collection.
func (rc *RockClient) AddDocumentsWithOffset(ctx context.Context, workspace, collection string,
	docs []interface{}) (openapi.AddDocumentsResponse, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.AddDocumentsResponse

	q := rc.DocumentsApi.AddDocuments(ctx, workspace, collection)

	// TODO should the method accept []map[string]interface{} instead?
	payload := make([]map[string]interface{}, len(docs))

	for i, d := range docs {
		payload[i] = d.(map[string]interface{})
	}

	req := openapi.NewAddDocumentsRequest(payload)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.AddDocumentsResponse{}, err
	}

	log := zerolog.Ctx(ctx)
	logDocumentStatuses(log.Trace(), resp.GetData()).Msg("added documents")

	return *resp, nil
}

// AddDocumentsRaw is a convenience method that adds new documents to a collection using a raw JSON payload,
// which has to be an array of objects.
//
// This method is useful when the documents are already in JSON format and do not need to be converted to
// a []map[string]interface{} slice, which is required by AddDocuments() and AddDocumentsWithOffset().
func (rc *RockClient) AddDocumentsRaw(ctx context.Context, workspace, collection string,
	docs json.RawMessage) ([]openapi.DocumentStatus, error) {
	return rawDocuments[*openapi.AddDocumentsResponse](ctx, rc, workspace, collection, http.MethodPost,
		io.MultiReader(bytes.NewBufferString(`{"data": `), bytes.NewReader(docs), bytes.NewBufferString("}")))
}

type patchDocumentRequest struct {
	Data []PatchDocument `json:"data"`
}

type PatchDocument struct {
	ID      string           `json:"_id"`
	Patches []PatchOperation `json:"patch"`
}

type PatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
	From  *string     `json:"from"`
}

// PatchDocuments patches (updates) existing documents in a collection. This convenience method does not use the
// generated client, as it takes values as map[string]interface{} which doesn't work when you want to patch e.g.
// a top-level boolean value.
//
// REST API documentation https://docs.rockset.com/rest-api/#patchdocuments
func (rc *RockClient) PatchDocuments(ctx context.Context, workspace, collection string,
	patches []PatchDocument) ([]openapi.DocumentStatus, error) {
	data := patchDocumentRequest{Data: patches}
	payload := bytes.Buffer{}
	e := json.NewEncoder(&payload)
	if err := e.Encode(data); err != nil {
		return nil, err
	}

	return rawDocuments[*openapi.PatchDocumentsResponse](ctx, rc, workspace, collection, http.MethodPatch, &payload)
}

type documentStatus interface {
	GetData() []openapi.DocumentStatus
}

// rawDocuments is a generic method to send raw documents to the Rockset REST API
func rawDocuments[T documentStatus](ctx context.Context, rc *RockClient, workspace, collection, method string,
	payload io.Reader) ([]openapi.DocumentStatus, error) {
	var err error
	var resp *http.Response

	req, err := rawDocumentsRequest(ctx, rc.RockConfig, workspace, collection, method, payload)
	if err != nil {
		return nil, err
	}

	err = rc.Retry(ctx, func() error {
		resp, err = rc.RockConfig.cfg.HTTPClient.Do(req)

		return rockerr.NewWithStatusCode(err, resp)
	})
	if err != nil {
		return nil, err
	}

	defer closeAndLog(ctx, resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log := zerolog.Ctx(ctx)

	// should this accept all 2xx status codes?
	if resp.StatusCode == http.StatusOK {
		var pdr T
		if err = json.Unmarshal(body, &pdr); err != nil {
			return nil, err
		}

		logDocumentStatuses(log.Trace(), pdr.GetData()).Msg("patched documents")

		return pdr.GetData(), nil
	}

	// if we get here, something went wrong, see if the error can be extracted from the body

	var em openapi.ErrorModel
	if err = json.Unmarshal(body, &em); err != nil {
		log.Error().Err(err).Str("body", string(body)).Msg("failed to unmarshal response")
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil, rockerr.Error{
		ErrorModel: &em,
		Cause:      fmt.Errorf("unexpected http response (%d): %s", resp.StatusCode, resp.Status),
	}
}

const docsURL = "https://%s/v1/orgs/self/ws/%s/collections/%s/docs"

func rawDocumentsRequest(ctx context.Context, rockConfig RockConfig, ws, collection, method string,
	payload io.Reader) (*http.Request, error) {
	// workspace and collection do not need to be escaped as they can only contain alphanumeric or dash characters
	u := fmt.Sprintf(docsURL, rockConfig.APIServer, ws, collection)
	req, err := http.NewRequestWithContext(ctx, method, u, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", rockConfig.cfg.UserAgent)

	// copy all default headers
	for k, v := range rockConfig.cfg.DefaultHeader {
		req.Header.Set(k, v)
	}

	return req, nil
}

func closeAndLog(ctx context.Context, c io.Closer) {
	if err := c.Close(); err != nil {
		log := zerolog.Ctx(ctx)
		log.Error().Err(err).Msg("failed to close")
	}
}

// DeleteDocuments deletes documents from a collection
//
// REST API documentation https://docs.rockset.com/rest-api/#deletedocuments
func (rc *RockClient) DeleteDocuments(ctx context.Context, workspace, collection string,
	docIDs []string) ([]openapi.DocumentStatus, error) {
	resp, err := rc.DeleteDocumentsWithOffset(ctx, workspace, collection, docIDs)
	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

// DeleteDocumentsWithOffset deletes documents from a collection, and returns the response with the offset(s),
// which can be used to wait until the collection has been updated to include the specified offset(s).
//
//	response, err := rs.DeleteDocumentsWithOffset(ctx, "commons", "users", docIDs)
//	if err != nil {
//	    return err
//	}
//	w := wait.New(rs)
//	err = w.UntilQueryable(ctx, "commons", "users", []string{response.GetLastOffset()})
//
// The added documents are now queryable in the collection.
func (rc *RockClient) DeleteDocumentsWithOffset(ctx context.Context, workspace, collection string,
	docIDs []string) (openapi.DeleteDocumentsResponse, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.DeleteDocumentsResponse

	q := rc.DocumentsApi.DeleteDocuments(ctx, workspace, collection)

	ids := make([]openapi.DeleteDocumentsRequestData, len(docIDs))
	for i, id := range docIDs {
		ids[i] = openapi.DeleteDocumentsRequestData{Id: id}
	}

	req := openapi.NewDeleteDocumentsRequest(ids)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.DeleteDocumentsResponse{}, err
	}

	log := zerolog.Ctx(ctx)
	logDocumentStatuses(log.Trace(), resp.GetData()).Msg("deleted documents")

	return *resp, nil
}

func logDocumentStatuses(e *zerolog.Event, statuses []openapi.DocumentStatus) *zerolog.Event {
	result := map[string]int{}
	for _, s := range statuses {
		i := result[s.GetStatus()]
		result[s.GetStatus()] = i + 1
	}
	for k, v := range result {
		e.Int(k, v)
	}

	return e
}
