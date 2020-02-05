package rockset_test

import (
	"fmt"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"log"
)

func Example_addDocument() {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	// document to be inserted
	m := map[string]interface{}{"name": "foo"}

	// array of documents
	docs := []interface{}{
		m,
	}

	dinfo := models.AddDocumentsRequest{
		Data: docs,
	}

	res, _, err := client.Documents.Add("commons", "test-collection", dinfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("document status: %s\n", res.Data[0].Status)
	// Output: document status: ADDED
}
