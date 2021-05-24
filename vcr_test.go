package rockset_test

import (
	"context"
	"fmt"
	"log"

	"github.com/seborama/govcr"

	"github.com/rockset/rockset-go-client"
)

// Use govcr to record API calls as fixtures and then replay them. Note that this will record your API key
// so do not commit the saved "cassette" to a public repo.
// See https://github.com/seborama/govcr for information how to use govcr.
func Example_vCR() {
	ctx := context.TODO()

	cfg := govcr.VCRConfig{Logging: true, CassettePath: "vcr"}
	rt := govcr.NewVCR("example", &cfg).Client

	rc, err := rockset.NewClient(rockset.FromEnv(), rockset.WithHTTPClient(rt))
	if err != nil {
		log.Fatal(err)
	}

	// first request will make a HTTP request to the Rockset API
	r, err := rc.GetOrganization(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("org: %s\n", r.GetId())

	// get a new client with recordings disabled
	cfg.DisableRecording = true
	rt = govcr.NewVCR("example", &cfg).Client
	rc, err = rockset.NewClient(rockset.FromEnv(), rockset.WithHTTPClient(rt))
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
