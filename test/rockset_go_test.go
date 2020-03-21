package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/antihax/optional"

	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	assert "github.com/stretchr/testify/require"
)

func TestCollection(t *testing.T) {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	workspace := "commons"
	name := "go-test-collection" + strconv.Itoa(rand.Intn(1000))

	{
		// create collection
		cinfo := models.CreateCollectionRequest{
			Name: name,
		}

		res, _, err := client.Collection.Create(workspace, cinfo)

		assert.Equal(t, err, nil, "error creating collection")
		assert.Equal(t, res.Data.Name, name, "collection should be created")
		assert.Equal(t, res.Data.Status, "CREATED", "collection status should be created")
	}

	{
		time.Sleep(5 * time.Second)

		// delete collection
		res, _, err := client.Collection.Delete(workspace, name)

		assert.Equal(t, err, nil, "error deleting collection")
		assert.Equal(t, res.Data.Name, name, "collection should be deleted")
		assert.Equal(t, res.Data.Status, "DELETED", "collection status should be deleted")
	}
}

func TestIntegration(t *testing.T) {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	name := "go-test-integration" + strconv.Itoa(rand.Intn(1000))

	{
		// create integration
		iinfo := models.CreateIntegrationRequest{
			Name: name,
			Dynamodb: &models.DynamodbIntegration{
				AwsAccessKey: &models.AwsAccessKey{
					AwsAccessKeyId:     ".....",
					AwsSecretAccessKey: ".....",
				},
			},
		}

		res, _, err := client.Integration.Create(iinfo)
		assert.Equal(t, err, nil, "error creating integration")
		assert.Equal(t, res.Data.Name, name, "integration should be created")
	}

	{
		// delete collection
		res, _, err := client.Integration.Delete(name)

		assert.Equal(t, err, nil, "error deleting integration")
		assert.Equal(t, res.Data.Name, name, "integration should be deleted")
	}
}

func TestQuery(t *testing.T) {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	{
		// construct query
		q := models.QueryRequest{
			Sql: &models.QueryRequestSql{
				Query: `select * from "_events" limit 1`,
			},
		}

		// query
		_, _, err := client.Query(q)
		assert.Equal(t, err, nil, "error querying")
	}
}

func TestQueryLambda(t *testing.T) {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	{
		// construct request
		q := models.CreateQueryLambdaRequest{
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
		_, _, err := client.QueryLambdas.Create("commons", q)
		assert.Equal(t, err, nil, "error creating Query Lambda")
	}

	{
		// execute Query Lambda with default paramaters
		res, _, err := client.QueryLambdas.Execute("commons", "MyQueryLambda", 1, nil)
		assert.Equal(t, err, nil, "error executing Query Lambda")

		results := res.Results[0].(map[string]interface{})

		assert.Equal(t, results["echo"].(string), "Hello, world!")
	}

	{
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

		res, _, err := client.QueryLambdas.Execute("commons", "MyQueryLambda", 1, &q)
		assert.Equal(t, err, nil, "error executing Query Lambda")

		results := res.Results[0].(map[string]interface{})

		assert.Equal(t, results["echo"].(string), "All work and no play makes Jack a dull boy")
	}

	{
		// delete Query Lambda
		_, _, err := client.QueryLambdas.Delete("commons", "MyQueryLambda")

		assert.Equal(t, err, nil, "error deleting Query Lambda")
	}
}
