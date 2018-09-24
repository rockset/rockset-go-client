package main

import (
	"fmt"

	apiclient "github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
)

func main() {
	// create the API client
	client := apiclient.Client("<api_key>", "<api_server>")

	// construct query
	q := models.QueryRequest{
		Sql: &models.QueryRequestSql{
			Query: "select * from \"_info.events\" where kind = :k limit 1",
			Parameters: []models.QueryParameter{
				{Name: "k", Type_: "string", Value: "QUERY"},
			},
		},
	}

	// query
	res, _, err := client.Query.Query(q, nil)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	// print result
	fmt.Println(res)
}
