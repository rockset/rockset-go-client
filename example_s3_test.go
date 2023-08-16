package rockset_test

import (
	"context"
	"fmt"
	"github.com/rockset/rockset-go-client/wait"
	"log"

	"github.com/rockset/rockset-go-client/option"
)

// Example code to first create an S3 integration, then create a collection from the integration, and finally clean up.
//
//nolint:funlen
func Example_s3() {
	ctx := context.TODO()

	rc, randomName, err := vcrClient("example_s3")
	if err != nil {
		log.Fatal(err)
	}

	integration := randomName("example_s3")
	collectionName := randomName("example_s3")
	workspace := "example"

	// create integration
	r, err := rc.CreateS3Integration(ctx, integration,
		option.AWSRole("arn:aws:iam::469279130686:role/rockset-s3-integration"),
		option.WithS3IntegrationDescription("created by go example code"))
	if err != nil {
		log.Fatalf("failed to create integration: %v", err)
	}
	fmt.Printf("integration created: %s\n", r.GetName())

	// create an S3 collection
	c, err := rc.CreateCollection(ctx, workspace, collectionName,
		option.WithCollectionDescription("created by go example code"),
		option.WithS3Source(integration, "rockset-go-tests",
			option.WithCSVFormat(
				[]string{"city", "country", "population", "visited"},
				[]option.ColumnType{
					option.ColumnTypeString, option.ColumnTypeString, option.ColumnTypeInteger, option.ColumnTypeBool,
				},
				option.WithEncoding("UTF-8"),
				option.WithEscapeChar("\\"),
				option.WithQuoteChar(`"`),
				option.WithSeparator(","),
			),
			option.WithS3Prefix("cities.csv"),
		),
		option.WithIngestTransformation("SELECT * FROM _input"),
	)
	if err != nil {
		log.Fatalf("failed to create collection: %v", err)
	}
	fmt.Printf("collection created: %s\n", c.GetName())

	// wait until collection is ready
	w := wait.New(rc)
	err = w.UntilCollectionReady(ctx, workspace, collectionName)
	if err != nil {
		log.Fatalf("failed to wait for collection to be ready: %v", err)
	}
	fmt.Printf("collection ready: %s\n", c.GetName())

	// wait until there are at least 3 new documents in the collection
	err = w.UntilCollectionHasDocuments(ctx, workspace, collectionName, 3)
	if err != nil {
		log.Fatalf("failed to wait for new documents: %v", err)
	}

	// get number of documents
	collection, err := rc.GetCollection(ctx, workspace, collectionName)
	if err != nil {
		log.Fatalf("failed to get collection: %v", err)
	}
	fmt.Printf("collection name: %s\n", collection.GetName())

	// delete the collection
	err = rc.DeleteCollection(ctx, workspace, collectionName)
	if err != nil {
		log.Fatalf("failed to delete collection: %v", err)
	}
	fmt.Printf("collection deleted: %s\n", c.GetName())

	// wait until the collection is gone
	err = w.UntilCollectionGone(ctx, workspace, collectionName)
	if err != nil {
		log.Fatalf("failed to wait for collection to be gone: %v", err)
	}
	fmt.Printf("collection gone: %s\n", c.GetName())

	// delete integration
	err = rc.DeleteIntegration(ctx, integration)
	if err != nil {
		log.Fatalf("failed to delete integration: %v", err)
	}
	fmt.Printf("integration deleted: %s\n", r.GetName())

	// Output:
	// integration created: go_example_s3
	// collection created: go_example_s3
	// collection ready: go_example_s3
	// collection name: go_example_s3
	// collection deleted: go_example_s3
	// collection gone: go_example_s3
	// integration deleted: go_example_s3
}
