package main

import (
	"fmt"

	apiclient "github.com/rockset/rockset-go-client"
)

func main() {
	// create the API client
	client := apiclient.Client("<api_key>", "<api_server>")

	// get collections
	res, _, err := client.Collection.Get("_info.events")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	res.PrintResponse()
}
