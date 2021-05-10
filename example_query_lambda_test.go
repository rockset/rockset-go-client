package rockset_test

import (
	"context"
	"fmt"
	"log"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

func ExampleRockClient_queryLambda() {
	ctx := context.TODO()

	rc, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.ExecuteQueryLambda(ctx, "commons", "test")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range *r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}

func ExampleRockClient_queryLambdaByTag() {
	ctx := context.TODO()

	rc, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.ExecuteQueryLambda(ctx, "commons", "test", option.WithTag("latest"))
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range *r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}

func ExampleRockClient_queryLambdaByVersion() {
	ctx := context.TODO()

	rc, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.ExecuteQueryLambda(ctx, "commons", "test", option.WithVersion("foobar"))
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range *r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}
