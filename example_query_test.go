package rockset_test

import (
	"fmt"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"log"
)

// This example runs a query against Rockset
func ExampleRockClient_Query() {
	// create the API client
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	// construct a query
	q := models.QueryRequest{
		Sql: &models.QueryRequestSql{
			Query: `SELECT * FROM "_events" LIMIT 1`,
		},
	}

	// execute the query
	res, _, err := client.Query(q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d row(s)\n", len(res.Results))
	// Output: got 1 row(s)
}
