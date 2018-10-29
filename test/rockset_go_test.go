package main

import (
	"math/rand"
	"strconv"
	"os"
	"testing"
	
	apiclient "github.com/rockset/rockset-go-client"
	assert "github.com/stretchr/testify/require"
	models "github.com/rockset/rockset-go-client/lib/go"
)

func TestCollection(t *testing.T) {
	apiKey := os.Getenv("ROCKSET_APIKEY")
	apiServer := os.Getenv("ROCKSET_APISERVER")

	client := apiclient.Client(apiKey, apiServer)
	
	name := "go-test-collection" + strconv.Itoa(rand.Intn(1000))

	{
		// create collection
		cinfo := models.CreateCollectionRequest{
			Name: name,
		}

		res, _, err := client.Collection.Create(cinfo)

		assert.Equal(t, err, nil, "error creating collection")
		assert.Equal(t, res.Data.Name, name, "collection should be created")
		assert.Equal(t, res.Data.Status, "CREATED", "collection status should be created")
	}

	{
		// delete collection
		res, _, err := client.Collection.Delete(name)

		assert.Equal(t, err, nil, "error deleting collection")
		assert.Equal(t, res.Data.Name, name, "collection should be deleted")
		assert.Equal(t, res.Data.Status, "DELETED", "collection status should be deleted")
	}	
}

func TestIntegration(t *testing.T) {
	apiKey := os.Getenv("ROCKSET_APIKEY")
	apiServer := os.Getenv("ROCKSET_APISERVER")

	client := apiclient.Client(apiKey, apiServer)
	
	name := "go-test-integration" + strconv.Itoa(rand.Intn(1000))

	{
		// create integration
		iinfo := models.CreateIntegrationRequest{
			Name: name,
			Aws: &models.AwsKeyIntegration{
				AwsAccessKeyId: ".....",
				AwsSecretAccessKey: ".....",
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
	apiKey := os.Getenv("ROCKSET_APIKEY")
	apiServer := os.Getenv("ROCKSET_APISERVER")

	client := apiclient.Client(apiKey, apiServer)
	
	{
		// construct query
		q := models.QueryRequest{
			Sql: &models.QueryRequestSql{
				Query: "select * from \"_events\" limit 1",
			},
		}

		// query
		_, _, err := client.Query(q)
		assert.Equal(t, err, nil, "error querying")
	}
}
