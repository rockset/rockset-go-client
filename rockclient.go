package rockset

import (
	"errors"
	"os"

	"github.com/rockset/rockset-go-client/openapi"
)

// Version is the Rockset client version
const Version = "0.11.0"

// DefaultAPIServer is the default Rockset API server to use
const DefaultAPIServer = "https://api.rs2.usw2.rockset.com"

const APIKeyEnvironmentVariableName = "ROCKSET_APIKEY"
const APIServerEnvironmentVariableName = "ROCKSET_APISERVER"

type RockClient struct {
	*openapi.APIClient
}

// NewClient creates a new Rockset client.
//
// Accessing the online database requires an API key, which you either have to supply through
// the ROCKSET_APIKEY environment variable and pass the FromEnv() option
//	c, err := rockset.NewClient(rockset.FromEnv())
// or explicitly using the WithAPIKey() option
//	c, err := rockset.NewClient(rockset.WithAPIKey("..."))
func NewClient(options ...RockOption) (*RockClient, error) {
	cfg := openapi.NewConfiguration()
	cfg.UserAgent = "rockset-go-client"
	cfg.AddDefaultHeader("x-rockset-version", Version)

	for _, o := range options {
		o(cfg)
	}

	// TODO should we pick up ROCKSET_APIKEY by default?
	c := openapi.NewAPIClient(cfg)
	return &RockClient{APIClient: c}, nil
}

// RockOption is the type for RockClient options
type RockOption func(rc *openapi.Configuration)

// FromEnv sets API key and API server from the environment variables ROCKSET_APIKEY and ROCKSET_APISERVER,
// and if ROCKSET_APISERVER is not set, it will use the default API server.
func FromEnv() RockOption {
	return func(cfg *openapi.Configuration) {
		if apikey, found := os.LookupEnv(APIKeyEnvironmentVariableName); found {
			cfg.AddDefaultHeader("Authorization", "apikey "+apikey)
		}

		if server, found := os.LookupEnv(APIServerEnvironmentVariableName); found {
			cfg.Host = server
		}
	}
}

func Debug() RockOption {
	return func(cfg *openapi.Configuration) {
		cfg.Debug = true
	}
}

func IsNotFoundError(err error) bool {
	var e openapi.GenericOpenAPIError
	if errors.As(err, &e) {
		return e.Error() == "404 Not Found"
	}
	return false
}
