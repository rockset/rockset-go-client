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

	rc, err := rockset.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.ExecuteQueryLambda(ctx, "commons", "eventType")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}

func ExampleRockClient_queryLambdaByTag() {
	ctx := context.TODO()

	rc, err := rockset.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.ExecuteQueryLambda(ctx, "commons", "eventType", option.WithTag("test"))
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}

func ExampleRockClient_queryLambdaByVersion() {
	ctx := context.TODO()

	rc, err := rockset.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.ExecuteQueryLambda(ctx, "commons", "eventType",
		option.WithVersion("e4e67e8835063d03"))
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}
