package rockset_test

import (
	"context"
	"fmt"
	"log"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

func ExampleRockClient_validateQuery() {
	ctx := context.TODO()

	rc, err := rockset.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r, err := rc.Query(ctx, "SELECT * FROM commons._events where label = :label",
		option.WithParameter("label", "string", "QUERY_SUCCESS"))
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range r.Collections {
		fmt.Printf("collection: %s\n", c)
	}

	// Output:
	// collection: commons._events
}
