package rockset_test

import (
	"context"
	"fmt"
	"log"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// Raw usage of the openapi client
func Example_queryRaw() {
	ctx := context.TODO()

	rc, err := rockset.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	q := rc.QueriesApi.Query(ctx)
	rq := openapi.NewQueryRequestWithDefaults()

	rq.Sql = openapi.QueryRequestSql{Query: "SELECT * FROM commons._events where label = :label"}
	rq.Sql.DefaultRowLimit = openapi.PtrInt32(10)

	rq.Sql.Parameters = []openapi.QueryParameter{
		{
			Name:  "label",
			Type:  "string",
			Value: "QUERY_SUCCESS",
		},
	}

	r, _, err := q.Body(*rq).Execute()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}

// Query convenience method
func ExampleRockClient_query() {
	ctx := context.TODO()

	rc, err := rockset.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.Query(ctx, "SELECT * FROM commons._events where label = :label",
		option.WithDefaultRowLimit(10),
		option.WithParameter("label", "string", "QUERY_SUCCESS"))
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}
