package main

import (
	"fmt"

	apiclient "github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
)

func main() {
	// create the API client
	client := apiclient.Client("<api_key>", "<api_server>")

	// create collection
	cinfo := models.CreateCollectionRequest{
		Name:        "my-first-collection",
		Description: "my first collection",
	}

	res, _, err := client.Collection.Create(cinfo)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	res.PrintResponse()
}
