package rockset_test

import (
	"context"
	"fmt"
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"log"
)

func ExampleRockClient_validateQuery() {
	ctx := context.TODO()

	rc, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.Query(ctx, "SELECT * FROM commons._events where label = :label",
		option.WithParameter("label", "string", "QUERY_SUCCESS"))
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range *r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}
