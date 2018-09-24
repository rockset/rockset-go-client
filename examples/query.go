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
			Query: "select * from \"_info.events\"",
		},
	}

	// query
	res, _, err := client.Query(q)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	// print result
	res.PrintResponse()
}
