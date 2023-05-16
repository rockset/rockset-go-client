package rockset_test

import (
	"context"
	"fmt"
	"log"

	"github.com/rockset/rockset-go-client"
	"github.com/seborama/govcr/v13"
)

// Use govcr to record API calls as fixtures and then replay them. The settings remove the HTTP header
// "Authorization" which is where the Rockset API key resides.
// See https://github.com/seborama/govcr for information how to use govcr.
func Example_vCR() {
	ctx := context.TODO()

	name := fmt.Sprintf("vcr/%s.cassette", "example_vcr")
	settings := vcrSettings(false)
	vcr := govcr.NewVCR(govcr.NewCassetteLoader(name), settings...)

	rc, err := rockset.NewClient(rockset.WithHTTPClient(vcr.HTTPClient()))
	if err != nil {
		log.Fatal(err)
	}

	// first request will make a HTTP request to the Rockset API
	r, err := rc.GetOrganization(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("org: %s\n", r.GetId())

	// get a new client with offline mode
	settings = vcrSettings(true)
	vcr = govcr.NewVCR(govcr.NewCassetteLoader(name), settings...)
	rc, err = rockset.NewClient(rockset.WithHTTPClient(vcr.HTTPClient()))
	if err != nil {
		log.Fatal(err)
	}

	// second request will replay the recorded information
	r, err = rc.GetOrganization(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("org: %s\n", r.GetId())

	// Output:
	// org: rockset-circleci
	// org: rockset-circleci
}
