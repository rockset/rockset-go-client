package rockset

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
)

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
	// APIKey is the API key to use for authentication
	APIKey string
	// APIServer is the API server to connect to
	APIServer string
	cfg       *openapi.Configuration
}

// RockClient is the client struct for making APi calls to Rockset.
type RockClient struct {
	*openapi.APIClient
	RockConfig
}

// NewClient creates a new Rockset client.
//
// Accessing the online database requires an API key and an API server to connect to,
// which are provided by through the ROCKSET_APIKEY and ROCKSET_APISERVER environment variables.
// If an API server isn't provided the DefaultAPIServer, is used.
//	c, err := rockset.NewClient()
// They can or explicitly using the WithAPIKey() and WithAPIServer() options.
//	c, err := rockset.NewClient(rockset.WithAPIKey("..."), rockset.WithAPIServer("..."))
func NewClient(options ...RockOption) (*RockClient, error) {
	cfg := openapi.NewConfiguration()
	cfg.UserAgent = "rockset-go-client"
	cfg.AddDefaultHeader("x-rockset-version", Version)
	// TODO should the default http client be tuned?
	cfg.HTTPClient = &http.Client{}

	rc := RockConfig{
		cfg:       cfg,
		Retrier:   ExponentialRetry{},
		APIKey:    os.Getenv(APIKeyEnvironmentVariableName),
		APIServer: os.Getenv(APIServerEnvironmentVariableName),
	}

	for _, o := range options {
		o(&rc)
	}

	if rc.APIServer == "" {
		rc.APIServer = DefaultAPIServer
	}

	if rc.APIKey == "" {
		return nil, errors.New("no API key provided")
	}
	cfg.AddDefaultHeader("Authorization", "apikey "+rc.APIKey)

	return &RockClient{
		RockConfig: rc,
		APIClient:  openapi.NewAPIClient(rc.cfg),
	}, nil
}

// RockOption is the type for RockClient options.
type RockOption func(rc *RockConfig)

// WithAPIKey sets the API key to use
func WithAPIKey(apiKey string) RockOption {
	return func(rc *RockConfig) {
		rc.cfg.AddDefaultHeader("Authorization", "apikey "+apiKey)
	}
}

// WithAPIServer sets the API server to connect to, and override the ROCKSET_APISERVER.
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

// WithHTTPDebug adds a http.RoundTripper that logs the request and response.
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
