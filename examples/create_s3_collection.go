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
		Name:        "s3_collection",
		Description: "s3 collection",
		Sources: []models.Source{
			{
				IntegrationName: "my-integration",
				S3: &models.SourceS3{
					Bucket: "<bucket>",
				},
			},
		},
	}

	rockresponse, _, err := client.Collection.Create(cinfo)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	rockresponse.PrintResponse()
}
