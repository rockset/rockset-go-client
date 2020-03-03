package rockset_test

import (
	"fmt"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"log"
	"time"
)

func Example_addDocument() {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	createResp, _, err := client.Collection.Create("commons", models.CreateCollectionRequest{
		Name:        "go-test-add-docs-collection",
		Description: "go test add docs collection",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created collection %s\n", createResp.Data.Name)

	time.Sleep(5 * time.Second)

	// document to be inserted
	m := map[string]interface{}{"name": "foo"}

	// array of documents
	docs := []interface{}{
		m,
	}

	dinfo := models.AddDocumentsRequest{
		Data: docs,
	}

	res, _, err := client.Documents.Add("commons", "go-test-add-docs-collection", dinfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("document status: %s\n", res.Data[0].Status)
	// Output: 
	// created collection go-test-add-docs-collection
	// document status: ADDED
	// deleted collection go-test-add-docs-collection

	deleteResp, _, err := client.Collection.Delete("commons", "go-test-add-docs-collection")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted collection %s\n", deleteResp.Data.Name)

}
