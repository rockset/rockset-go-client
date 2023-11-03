package paginate

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

const DefaultPaginatedQuerySize = 1000

type Paginator struct {
	rc RockClient
	// Max number of documents to get. Defaults to DefaultPaginatedQuerySize, and must be a value less than 100,000
	PageSize int
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fake . RockClient
type RockClient interface {
	Query(ctx context.Context, sql string, options ...option.QueryOption) (openapi.QueryResponse, error)
	GetQueryResults(ctx context.Context, queryID string, options ...option.QueryResultOption) (openapi.QueryPaginationResponse, error)
}

// New creates a new Paginator with PageSize set to DefaultPaginatedQuerySize
func New(rc RockClient) *Paginator {
	return &Paginator{rc, DefaultPaginatedQuerySize}
}

// Query executes a query and blocks until it has retrieved all results using pagination
// and sent them to the document channel, which is closed when all documents have been sent.
// This sets option.WithMaxInitialResults(DefaultPaginatedQuerySize),
// which is also used to fetch the paginated results using the GetQueryResults(),
// but it can be overridden by setting the PageSize on the Paginator.
func (p Paginator) Query(ctx context.Context, documentCh chan<- map[string]any, sql string, options ...option.QueryOption) error {
	defer close(documentCh)

	options = append(options, option.WithMaxInitialResults(int64(p.PageSize)))

	response, err := p.rc.Query(ctx, sql, options...)
	if err != nil {
		return err
	}

	for _, doc := range response.Results {
		documentCh <- doc
	}
	cursor := response.Pagination.GetNextCursor()

	for {
		if cursor == "" {
			break
		}

		// get next batch of query results
		res, err := p.rc.GetQueryResults(ctx, response.GetQueryId(), option.WithQueryResultCursor(cursor),
			option.WithQueryResultDocs(int32(p.PageSize)))
		if err != nil {
			return err
		}

		// send documents from the current batch
		for _, doc := range res.Results {
			documentCh <- doc
		}
		cursor = res.Pagination.GetNextCursor()
	}

	return nil
}

// GetQueryResults gets all query results from the previously executed query with queryID,
// and sends the result to the document channel, which is closed once everything has been retrieved.
func (p Paginator) GetQueryResults(ctx context.Context, documentCh chan<- map[string]any, queryID string) error {
	defer close(documentCh)

	var cursor string
	for {
		var options []option.QueryResultOption
		if cursor != "" {
			options = append(options, option.WithQueryResultCursor(cursor))
		}

		res, err := p.rc.GetQueryResults(ctx, queryID, options...)
		if err != nil {
			return err
		}

		// TODO if the query hasn't finished running, this could optionally go into wait loop

		// send documents from the current batch
		for _, doc := range res.Results {
			documentCh <- doc
		}
		cursor = res.Pagination.GetNextCursor()

		if cursor == "" {
			break
		}
	}

	return nil
}
