package paginate_test

import (
	"errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/paginate/fake"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/paginate"
)

func TestQueryPaginated_integration(t *testing.T) {
	test.SkipUnlessIntegrationTest(t)
	ctx := test.Context()
	rc := test.Client(t)
	p := paginate.New(rc)
	p.PageSize = 500

	docs := make(chan map[string]any, p.PageSize)

	wg := sync.WaitGroup{}
	wg.Add(1)
	var count int
	go func() {
		for range docs {
			count++
		}
		wg.Done()
	}()

	err := p.Query(ctx, docs, "SELECT * FROM persistent.snp")
	require.NoError(t, err)

	wg.Wait()
	assert.Equal(t, 2632, count)
}

func TestQueryPaginated(t *testing.T) {
	ctx := test.Context()

	rc := &fake.FakeRockClient{}
	rc.QueryReturns(openapi.QueryResponse{
		Results: []map[string]interface{}{
			{"0": "0"},
			{"1": "1"},
		},
		Pagination: &openapi.PaginationInfo{NextCursor: openapi.PtrString("foobar")},
	}, nil)
	rc.GetQueryResultsReturns(openapi.QueryPaginationResponse{
		Results: []map[string]interface{}{
			{"2": "2"},
			{"3": "3"},
		},
	}, nil)

	p := paginate.New(rc)

	docs := make(chan map[string]any, 100)

	wg := sync.WaitGroup{}
	wg.Add(1)
	var count int
	go func() {
		for range docs {
			count++
		}
		wg.Done()
	}()

	err := p.Query(ctx, docs, "SELECT * FROM persistent.snp")
	require.NoError(t, err)

	wg.Wait()
	assert.Equal(t, 4, count)
	assert.Equal(t, 1, rc.QueryCallCount())
	assert.Equal(t, 1, rc.GetQueryResultsCallCount())
}

func TestQueryPaginated_queryError(t *testing.T) {
	ctx := test.Context()

	rc := &fake.FakeRockClient{}
	rc.QueryReturns(openapi.QueryResponse{}, errors.New("boom"))

	p := paginate.New(rc)

	docs := make(chan map[string]any, 100)

	wg := sync.WaitGroup{}
	wg.Add(1)
	var count int
	go func() {
		for range docs {
			count++
		}
		wg.Done()
	}()

	err := p.Query(ctx, docs, "SELECT * FROM persistent.snp")
	assert.Error(t, err)

	wg.Wait()
	assert.Equal(t, 0, count)
	assert.Equal(t, 1, rc.QueryCallCount())
	assert.Equal(t, 0, rc.GetQueryResultsCallCount())
}

func TestQueryPaginated_paginationError(t *testing.T) {
	ctx := test.Context()

	rc := &fake.FakeRockClient{}
	rc.QueryReturns(openapi.QueryResponse{}, errors.New("boom"))
	rc.QueryReturns(openapi.QueryResponse{
		Results: []map[string]interface{}{
			{"0": "0"},
			{"1": "1"},
		},
		Pagination: &openapi.PaginationInfo{NextCursor: openapi.PtrString("foobar")},
	}, nil)
	rc.GetQueryResultsReturns(openapi.QueryPaginationResponse{}, errors.New("boom"))

	p := paginate.New(rc)

	docs := make(chan map[string]any, 100)

	wg := sync.WaitGroup{}
	wg.Add(1)
	var count int
	go func() {
		for range docs {
			count++
		}
		wg.Done()
	}()

	err := p.Query(ctx, docs, "SELECT * FROM persistent.snp")
	assert.Error(t, err)

	wg.Wait()
	assert.Equal(t, 2, count)
	assert.Equal(t, 1, rc.QueryCallCount())
	assert.Equal(t, 1, rc.GetQueryResultsCallCount())
}

func TestPaginator_GetQueryResults(t *testing.T) {
	ctx := test.Context()

	rc := &fake.FakeRockClient{}
	rc.GetQueryResultsReturnsOnCall(0, openapi.QueryPaginationResponse{
		Pagination: &openapi.PaginationInfo{
			NextCursor: openapi.PtrString("foo"),
		},
		Results: []map[string]interface{}{
			{"0": "0", "1": "1"},
		},
	}, nil)
	rc.GetQueryResultsReturnsOnCall(1, openapi.QueryPaginationResponse{
		Pagination: &openapi.PaginationInfo{
			NextCursor: openapi.PtrString(""),
		},
		Results: []map[string]interface{}{
			{"2": "2", "3": "3"},
		},
	}, nil)
	docs := make(chan map[string]any, 100)

	wg := sync.WaitGroup{}
	wg.Add(1)
	var count int
	go func() {
		for range docs {
			count++
		}
		wg.Done()
	}()

	p := paginate.New(rc)

	err := p.GetQueryResults(ctx, docs, "id")
	assert.NoError(t, err)

	wg.Wait()
	assert.Equal(t, 2, count)
	assert.Equal(t, 2, rc.GetQueryResultsCallCount())
}
