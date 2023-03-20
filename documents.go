package rockset

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
)

// https://docs.rockset.com/rest-api/#documents

// AddDocuments adds new documents to a collection
//
// REST API documentation https://docs.rockset.com/rest-api/#adddocuments
func (rc *RockClient) AddDocuments(ctx context.Context, workspace, collection string,
	docs []interface{}) ([]openapi.DocumentStatus, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.AddDocumentsResponse

	log := zerolog.Ctx(ctx)
	q := rc.DocumentsApi.AddDocuments(ctx, workspace, collection)

	// TODO should the method accept []map[string]interface{} instead?
	payload := make([]map[string]interface{}, len(docs))

	for i, d := range docs {
		payload[i] = d.(map[string]interface{})
	}

	req := openapi.NewAddDocumentsRequest(payload)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return NewErrorWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	logDocumentStatuses(log.Trace(), resp.GetData()).Msg("added documents")

	return resp.GetData(), nil
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
	var err error
	var resp *http.Response

	req, err := rc.patchDocumentsRequest(ctx, workspace, collection, patches)
	if err != nil {
		return nil, err
	}

	err = rc.Retry(ctx, func() error {
		resp, err = rc.RockConfig.cfg.HTTPClient.Do(req)

		return NewErrorWithStatusCode(err, resp)
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
		var pdr openapi.PatchDocumentsResponse
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

	return nil, Error{
		ErrorModel: &em,
		Cause:      fmt.Errorf("unexpected http response (%d): %s", resp.StatusCode, resp.Status),
	}
}

func (rc *RockClient) patchDocumentsRequest(ctx context.Context, ws, collection string,
	patches []PatchDocument) (*http.Request, error) {
	data := patchDocumentRequest{Data: patches}
	payload := bytes.Buffer{}
	e := json.NewEncoder(&payload)
	if err := e.Encode(data); err != nil {
		return nil, err
	}

	// workspace and collection do not need to be escaped as they can only contain alphanumeric or dash characters
	u := fmt.Sprintf("https://%s/v1/orgs/self/ws/%s/collections/%s/docs", rc.RockConfig.APIServer, ws, collection)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, u, &payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "apikey "+rc.RockConfig.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-rockset-version", Version)
	req.Header.Set("User-Agent", rc.RockConfig.cfg.UserAgent)

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
	var err error
	var httpResp *http.Response
	var resp *openapi.DeleteDocumentsResponse

	log := zerolog.Ctx(ctx)
	q := rc.DocumentsApi.DeleteDocuments(ctx, workspace, collection)

	ids := make([]openapi.DeleteDocumentsRequestData, len(docIDs))
	for i, id := range docIDs {
		ids[i] = openapi.DeleteDocumentsRequestData{Id: id}
	}

	req := openapi.NewDeleteDocumentsRequest(ids)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = q.Body(*req).Execute()

		return NewErrorWithStatusCode(err, httpResp)
	})

	if err != nil {
		return nil, err
	}

	logDocumentStatuses(log.Trace(), resp.GetData()).Msg("deleted documents")

	return resp.GetData(), nil
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
