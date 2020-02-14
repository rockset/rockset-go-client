package rockset_test

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"log"
)

// Example code to create, use, and delete a saved query
func Example_savedQuery() {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	var defaultValue interface{} = "Hello, world!"

	// construct request
	createQuery := models.CreateSavedQueryRequest{
		Name:     "MySavedQuery",
		QuerySql: `SELECT :param AS echo`,
		Parameters: []models.SavedQueryParameter{
			{
				Name:         "param",
				Type_:        "string",
				DefaultValue: &defaultValue,
			},
		},
	}

	// create saved query
	createResp, _, err := client.QueryApi.Create("commons", createQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created saved query %s\n", createResp.Data.Name)

	// execute saved query with default parameters
	runResp, _, err := client.QueryApi.Run("commons", "MySavedQuery", 1, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("query result: %v\n", runResp.Results[0])

	// execute saved query with explicit parameters
	var customValue interface{} = "All work and no play makes Jack a dull boy"
	q := models.RunOpts{
		Body: optional.NewInterface(models.ExecuteSavedQueryRequest{
			Parameters: []models.ExecuteSavedQueryParameter{
				{
					Name:  "param",
					Value: &customValue,
				},
			},
		}),
	}

	runResp, _, err = client.QueryApi.Run("commons", "MySavedQuery", 1, &q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("query result: %v\n", runResp.Results[0])

	// delete saved query
	deleteResp, _, err := client.QueryApi.Delete("commons", "MySavedQuery")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted saved query %s\n", deleteResp.Data.Name)

	// Output:
	// created saved query MySavedQuery
	// query result: map[echo:Hello, world!]
	// query result: map[echo:All work and no play makes Jack a dull boy]
	// deleted saved query MySavedQuery
}
