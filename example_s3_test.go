package rockset_test

import (
	"context"
	"fmt"
	"log"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

// Example code to first create an S3 integration, then create a collection from the integration, and finally clean up.
//
//nolint:funlen
func Example_s3() {
	ctx := context.TODO()

	rc, err := rockset.NewClient(rockset.WithHTTPDebug())
	if err != nil {
		log.Fatal(err)
	}

	// create integration
	r, err := rc.CreateS3Integration(ctx, "s3exampleIntegration",
		option.AWSRole("arn:aws:iam::469279130686:role/rockset-s3-integration"),
		option.WithS3IntegrationDescription("created by go example code"))
	if err != nil {
		log.Fatalf("failed to create integration: %v", err)
	}
	fmt.Printf("integration created: %s\n", r.GetName())

	// create an S3 collection
	c, err := rc.CreateCollection(ctx, "commons", "s3example",
		option.WithCollectionDescription("created by go example code"),
		option.WithS3Source("s3exampleIntegration", "rockset-go-tests",
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
		option.WithInsertOnly(),
		option.WithFieldMappingQuery("SELECT * FROM _input"),
	)
	if err != nil {
		log.Fatalf("failed to create collection: %v", err)
	}
	fmt.Printf("collection created: %s\n", c.GetName())

	// wait until collection is ready
	err = rc.WaitUntilCollectionReady(ctx, "commons", "s3example")
	if err != nil {
		log.Fatalf("failed to wait for collection to be ready: %v", err)
	}
	fmt.Printf("collection ready: %s\n", c.GetName())

	// wait until there are at least 3 new documents in the collection
	err = rc.WaitUntilCollectionDocuments(ctx, "commons", "s3example", 3)
	if err != nil {
		log.Fatalf("failed to wait for new documents: %v", err)
	}

	// get number of documents
	collection, err := rc.GetCollection(ctx, "commons", "s3example")
	if err != nil {
		log.Fatalf("failed to get collection: %v", err)
	}
	fmt.Printf("collection documents: %d\n", collection.Stats.GetDocCount())

	// delete the collection
	err = rc.DeleteCollection(ctx, "commons", "s3example")
	if err != nil {
		log.Fatalf("failed to delete collection: %v", err)
	}
	fmt.Printf("collection deleted: %s\n", c.GetName())

	// wait until the collection is gone
	err = rc.WaitUntilCollectionGone(ctx, "commons", "s3example")
	if err != nil {
		log.Fatalf("failed to wait for collection to be gone: %v", err)
	}
	fmt.Printf("collection gone: %s\n", c.GetName())

	// delete integration
	err = rc.DeleteIntegration(ctx, "s3exampleIntegration")
	if err != nil {
		log.Fatalf("failed to delete integration: %v", err)
	}
	fmt.Printf("integration deleted: %s\n", r.GetName())

	// Output:
	// integration created: s3exampleIntegration
	// collection created: s3example
	// collection ready: s3example
	// collection documents: 3
	// collection deleted: s3example
	// collection gone: s3example
	// integration deleted: s3exampleIntegration
}
