package rockset_test

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"log"
)

// Example code to create, use, and delete a Query Lambda
func Example_queryLambda() {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	// construct request
	createLambda := models.CreateQueryLambdaRequest{
		Name:     "MyQueryLambda",
		Sql: &models.QueryLambdaSql{
			Query: "SELECT :param as echo",
			DefaultParameters: []models.QueryParameter{
				{
					Name:  "param",
					Type_: "string",
					Value: "Hello, world!",
				},
			},
		},
	}

	// create Query Lambda
	createResp, _, err := client.QueryLambdas.Create("commons", createLambda)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created Query Lambda %s\n", createResp.Data.Name)

	// execute Query Lambda with default parameters
	runResp, _, err := client.QueryLambdas.Execute("commons", "MyQueryLambda", 1, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("query result: %v\n", runResp.Results[0])

	// execute Query Lambda with explicit paramaters
	q := models.ExecuteOpts{
		Body: optional.NewInterface(models.ExecuteQueryLambdaRequest{
			Parameters: []models.QueryParameter{
				{
					Name:  "param",
					Type_: "string",
					Value: "All work and no play makes Jack a dull boy",
				},
			},
		}),
	}

	runResp, _, err = client.QueryLambdas.Execute("commons", "MyQueryLambda", 1, &q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("query result: %v\n", runResp.Results[0])

	deleteResp, _, err := client.QueryLambdas.Delete("commons", "MyQueryLambda")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted Query Lambda %s\n", deleteResp.Data.Name)

	// Output:
	// created Query Lambda MyQueryLambda
	// query result: map[echo:Hello, world!]
	// query result: map[echo:All work and no play makes Jack a dull boy]
	// deleted Query Lambda MyQueryLambda
}
