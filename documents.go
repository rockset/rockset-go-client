package rockset

import (
	"context"

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
	var resp openapi.AddDocumentsResponse

	log := zerolog.Ctx(ctx)
	q := rc.DocumentsApi.AddDocuments(ctx, workspace, collection)

	// TODO should the method accept []map[string]interface{} instead?
	payload := make([]map[string]interface{}, len(docs))

	for i, d := range docs {
		payload[i] = d.(map[string]interface{})
	}

	req := openapi.NewAddDocumentsRequest(payload)

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	logResult(log.Trace(), *resp.Data).Msg("added documents")

	return *resp.Data, nil
}

// PatchDocuments updates (patches) existing documents in a collection
//
// REST API documentation https://docs.rockset.com/rest-api/#patchdocuments
func (rc *RockClient) PatchDocuments(ctx context.Context, workspace, collection string,
	docs []openapi.PatchDocument) ([]openapi.DocumentStatus, error) {
	var err error
	var resp openapi.PatchDocumentsResponse

	log := zerolog.Ctx(ctx)
	q := rc.DocumentsApi.PatchDocuments(ctx, workspace, collection)

	req := openapi.NewPatchDocumentsRequest(docs)

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return resp.Data, err
	}

	logResult(log.Trace(), resp.Data).Msg("patched documents")

	return resp.Data, nil
}

// DeleteDocuments deletes documents from a collection
//
// REST API documentation https://docs.rockset.com/rest-api/#deletedocuments
func (rc *RockClient) DeleteDocuments(ctx context.Context, workspace, collection string,
	docIDs []string) ([]openapi.DocumentStatus, error) {
	var err error
	var resp openapi.DeleteDocumentsResponse

	log := zerolog.Ctx(ctx)
	q := rc.DocumentsApi.DeleteDocuments(ctx, workspace, collection)

	ids := make([]openapi.DeleteDocumentsRequestData, len(docIDs))
	for i, id := range docIDs {
		ids[i] = openapi.DeleteDocumentsRequestData{Id: id}
	}

	req := openapi.NewDeleteDocumentsRequest(ids)

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	logResult(log.Trace(), *resp.Data).Msg("deleted documents")

	return *resp.Data, nil
}

func logResult(e *zerolog.Event, statuses []openapi.DocumentStatus) *zerolog.Event {
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
