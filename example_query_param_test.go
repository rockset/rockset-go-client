package rockset_test

import (
	"fmt"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"log"
)

func ExampleRockClient_queryParam() {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	// construct query
	q := models.QueryRequest{
		Sql: &models.QueryRequestSql{
			Query: `SELECT :k`,
			Parameters: []models.QueryParameter{
				{Name: "k", Type_: "string", Value: "foo"},
			},
		},
	}

	// execute query
	res, _, err := client.Query(q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d response(s)\n", len(res.Results))
	// Output: got 1 response(s)
}
