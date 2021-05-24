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

	err = rc.Retry(ctx, func() (bool, error) {
		resp, _, err = q.Body(*req).Execute()
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}
			return false, re
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	// TODO add based on status: ADDED, REPLACED, DELETED, ERROR
	log.Trace().Int("count", len(*resp.Data)).Msg("added documents")

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

	err = rc.Retry(ctx, func() (bool, error) {
		resp, _, err = q.Body(*req).Execute()
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}
			return false, re
		}

		return false, nil
	})

	if err != nil {
		return resp.Data, err
	}

	// TODO add based on status: ADDED, REPLACED, DELETED, ERROR
	log.Trace().Int("count", len(resp.Data)).Msg("patched documents")

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

	err = rc.Retry(ctx, func() (bool, error) {
		resp, _, err = q.Body(*req).Execute()
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}
			return false, re
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	// TODO add based on status: ADDED, REPLACED, DELETED, ERROR
	log.Trace().Int("count", len(*resp.Data)).Msg("deleted documents")

	return *resp.Data, nil
}
