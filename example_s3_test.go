package rockset_test

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// Example code to first create an S3 integration, then create a collection from the integration,
// and finally clean up.
//nolint:funlen
func Example_s3() {
	ctx := context.TODO()

	rc, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	// create integration
	r, err := rc.CreateS3Integration(ctx, "s3exampleIntegration",
		option.AWSRole("arn:aws:iam::216690786812:role/rockset-integration-role"),
		option.WithS3IntegrationDescription("created by go example code"))
	if err != nil {
		log.Fatalf("failed to create integration: %v", err)
	}
	fmt.Printf("integration created: %s\n", r.Data.GetName())

	// create collection
	c, err := rc.CreateS3Collection(ctx, "commons", "s3example", "created by go example code",
		"s3exampleIntegration", "rockset-terraform-provider", "cities.csv",
		rockset.WithCSVFormat(
			[]string{"city", "country", "population", "visited"},
			[]rockset.ColumnType{
				rockset.ColumnTypeSTRING, rockset.ColumnTypeSTRING, rockset.ColumnTypeInteger, rockset.ColumnTypeBool,
			},
			option.WithEncoding("UTF-8"),
			option.WithEscapeChar("\\"),
			option.WithQuoteChar(`"`),
			option.WithSeparator(","),
		),
		option.WithCollectionFieldSchema("city", option.WithColumnIndexMode(option.ColumnIndexModeNoStore)),
		option.WithCollectionFieldMapping("test", false,
			option.OutputField("out", "CAST(:country AS string)", option.OnErrorSkip),
			option.InputField("country", option.FieldMissingSkip, false, "country")),
	)
	if err != nil {
		var e openapi.GenericOpenAPIError
		if errors.As(err, &e) {
			log.Printf("err: %s", string(e.Body()))
		}
		log.Fatalf("failed to create collection: %v", err)
	}
	fmt.Printf("collection created: %s\n", c.GetName())

	// wait until collection is ready
	err = rc.WaitForCollectionReady(ctx, "commons", "s3example")
	if err != nil {
		log.Fatalf("failed to wait for collection to be ready: %v", err)
	}
	fmt.Printf("collection ready: %s\n", c.GetName())

	// wait until there are at least 3 new documents in the collection
	err = rc.WaitForCollectionDocuments(ctx, "commons", "s3example", 3)
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
	err = rc.WaitForCollectionGone(ctx, "commons", "s3example")
	if err != nil {
		log.Fatalf("failed to wait for collection to be gone: %v", err)
	}
	fmt.Printf("collection gone: %s\n", c.GetName())

	// delete integration
	err = rc.DeleteIntegration(ctx, "s3exampleIntegration")
	if err != nil {
		log.Fatalf("failed to delete integration: %v", err)
	}
	fmt.Printf("integration deleted: %s\n", r.Data.GetName())

	// Output:
	// integration created: s3exampleIntegration
	// collection created: s3example
	// collection ready: s3example
	// collection documents: 3
	// collection deleted: s3example
	// collection gone: s3example
	// integration deleted: s3exampleIntegration
}
