package rockset

import (
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
)

// Version is the Rockset client version
const Version = "0.11.0"

// DefaultAPIServer is the default Rockset API server to use
const DefaultAPIServer = "https://api.rs2.usw2.rockset.com"

// APIKeyEnvironmentVariableName is the environment variable name for the API key
const APIKeyEnvironmentVariableName = "ROCKSET_APIKEY"

// APIServerEnvironmentVariableName is the environment variable name for the API server
const APIServerEnvironmentVariableName = "ROCKSET_APISERVER"

// RockConfig contains the configurable options for the RockClient.
type RockConfig struct {
	// Retrier is the retry function used to retry API calls.
	Retrier
	cfg *openapi.Configuration
}

// RockClient is the client struct for making APi calls to Rockset.
type RockClient struct {
	*openapi.APIClient
	RockConfig
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
	// TODO should the default http client be tuned?
	cfg.HTTPClient = &http.Client{}

	// set defaults
	cfg.Host = os.Getenv(APIServerEnvironmentVariableName)
	cfg.AddDefaultHeader("Authorization", "apikey "+os.Getenv(APIKeyEnvironmentVariableName))

	rc := RockConfig{
		cfg:     cfg,
		Retrier: ExponentialRetry{},
	}

	for _, o := range options {
		o(&rc)
	}

	return &RockClient{
		RockConfig: rc,
		APIClient:  openapi.NewAPIClient(rc.cfg),
	}, nil
}

// RockOption is the type for RockClient options
type RockOption func(rc *RockConfig)

// FromEnv sets API key and API server from the environment variables ROCKSET_APIKEY and ROCKSET_APISERVER,
// and if ROCKSET_APISERVER is not set, it will use the default API server.
func FromEnv() RockOption {
	return func(rc *RockConfig) {
		if apikey, found := os.LookupEnv(APIKeyEnvironmentVariableName); found {
			rc.cfg.AddDefaultHeader("Authorization", "apikey "+apikey)
		}

		if server, found := os.LookupEnv(APIServerEnvironmentVariableName); found {
			rc.cfg.Host = server
		}
	}
}

// WithAPIKey sets the API key to use
func WithAPIKey(apiKey string) RockOption {
	return func(rc *RockConfig) {
		rc.cfg.AddDefaultHeader("Authorization", "apikey "+apiKey)
	}
}

// WithAPIServer sets the API server to connect to
func WithAPIServer(server string) RockOption {
	return func(rc *RockConfig) {
		rc.cfg.Host = server
	}
}

// WithHTTPClient sets the HTTP client. Without this option RockClient uses the http.DefaultClient.
func WithHTTPClient(c *http.Client) RockOption {
	return func(rc *RockConfig) {
		rc.cfg.HTTPClient = c
	}
}

// WithRetry sets the Retrier the RockClient uses to retry requests which return a Error that can be retried.
func WithRetry(r Retrier) RockOption {
	return func(rc *RockConfig) {
		rc.Retrier = r
	}
}

// WithHTTPDebug adds a http.RoundTripper that logs the request and response
func WithHTTPDebug() RockOption {
	return func(rc *RockConfig) {
		rc.cfg.HTTPClient.Transport = &debugRoundTripper{http.DefaultTransport}
	}
}

type debugRoundTripper struct {
	transport http.RoundTripper
}

func (r *debugRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	log := zerolog.Ctx(ctx)

	reqb, _ := httputil.DumpRequest(req, true)
	res, err := r.transport.RoundTrip(req)
	resb, _ := httputil.DumpResponse(res, true)

	// TODO create a custom dump as this contains the api key
	log.Debug().Str("data", string(reqb)).Msg("request")
	log.Debug().Str("data", string(resb)).Msg("response")

	return res, err
}
