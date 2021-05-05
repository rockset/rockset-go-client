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
		log.Fatalf("failed to create Rockset client: %v", err)
	}

	fatal := func(msg string, err error) {
		if e, ok := rockset.AsRocksetError(err); ok {
			log.Fatalf("%s (%s): %v", msg, e.Message, err)
		}
		log.Fatalf("%s: %v", msg, err)
	}

	workspace := "commons"
	lambdaName := "MyQueryLambda"

	// construct request
	createLambda := models.CreateQueryLambdaRequest{
		Name: lambdaName,
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
	createResp, _, err := client.QueryLambdas.Create(workspace, createLambda)
	if err != nil {
		fatal("failed to create query lambda", err)
	}

	fmt.Printf("created Query Lambda %s\n", createResp.Data.Name)

	// execute Query Lambda with default parameters
	runResp, _, err := client.QueryLambdas.Execute_4(workspace, lambdaName, createResp.Data.Version, nil)
	if err != nil {
		fatal("failed to execute query lambda", err)
	}

	fmt.Printf("query result: %v\n", runResp.Results[0])

	// execute Query Lambda with explicit parameters
	q := models.Execute_4Opts{
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

	runResp, _, err = client.QueryLambdas.Execute_4(workspace, lambdaName, createResp.Data.Version, &q)
	if err != nil {
		fatal("failed to execute query lambda with explicit params", err)
	}

	fmt.Printf("query result: %v\n", runResp.Results[0])

	deleteResp, _, err := client.QueryLambdas.Delete(workspace, lambdaName)
	if err != nil {
		fatal("failed to delete query lambda", err)
	}

	fmt.Printf("deleted Query Lambda %s\n", deleteResp.Data.Name)

	// Output:
	// created Query Lambda MyQueryLambda
	// query result: map[echo:Hello, world!]
	// query result: map[echo:All work and no play makes Jack a dull boy]
	// deleted Query Lambda MyQueryLambda
}
