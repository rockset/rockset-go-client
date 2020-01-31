package main

import (
	"fmt"
	"os"

	"github.com/antihax/optional"
	apiclient "github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
)

func main() {
	apiKey := os.Getenv("ROCKSET_APIKEY")
	apiServer := os.Getenv("ROCKSET_APISERVER")

	// create the API client
	client := apiclient.Client(apiKey, apiServer)

	var defaultValue interface{} = "Hello, world!"

	{
		// construct request
		q := models.CreateSavedQueryRequest{
			Name:     "MySavedQuery",
			QuerySql: "SELECT :param as echo",
			Parameters: []models.SavedQueryParameter{
				{
					Name:         "param",
					Type_:        "string",
					DefaultValue: &defaultValue,
				},
			},
		}

		// create saved query
		res, _, err := client.QueryApi.Create("commons", q)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		res.PrintResponse()
	}

	{
		// execute saved query with default paramaters
		res, _, err := client.QueryApi.Run("commons", "MySavedQuery", 1, nil)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		res.PrintResponse()
	}

	{
		// execute saved query with explicit paramaters
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

		res, _, err := client.QueryApi.Run("commons", "MySavedQuery", 1, &q)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		res.PrintResponse()
	}

	{
		// delete saved query
		res, _, err := client.QueryApi.Delete("commons", "MySavedQuery")

		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		res.PrintResponse()
	}
}
