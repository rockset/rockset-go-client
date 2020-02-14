package rockset_test

import (
	"fmt"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"log"
)

func Example_createCollection() {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	// create collection
	createResp, _, err := client.Collection.Create("commons", models.CreateCollectionRequest{
		Name:        "go-test-collection",
		Description: "go test collection",
	})
	fmt.Printf("created collection %s\n", createResp.Data.Name)

	// get collection
	getResp, _, err := client.Collection.Get("commons", "go-test-collection")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("collection status for %s: %s\n", getResp.Data.Name, getResp.Data.Status)

	// delete collection
	deleteResp, _, err := client.Collection.Delete("commons", "go-test-collection")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted collection %s\n", deleteResp.Data.Name)

	// Output:
	// created collection go-test-collection
	// collection status for go-test-collection: CREATED
	// deleted collection go-test-collection
}

func Example_createS3Collection() {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	// create collection
	req := models.CreateCollectionRequest{
		Name:        "s3_collection",
		Description: "s3 collection",
		Sources: []models.Source{
			{
				IntegrationName: "my-first-integration",
				S3: &models.SourceS3{
					Bucket: "s3://<bucket-name>",
				},
			},
		},
	}

	resp, _, err := client.Collection.Create("commons", req)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	resp.PrintResponse()
}
