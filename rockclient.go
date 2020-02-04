package rockset

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	api "github.com/rockset/rockset-go-client/lib/go"
)

// Version is the Rockset client version
const Version = "0.8.0"

// DefaultAPIServer is the default Rockset API server to use
const DefaultAPIServer = "https://api.rs2.usw2.rockset.com"

type RockClient struct {
	apiServer  string
	apiKey     string
	timeout    time.Duration
	httpClient *http.Client

	common api.Service

	// API Services
	ApiKeys     *api.ApiKeysApiService
	Collection  *api.CollectionsApiService
	Integration *api.IntegrationsApiService
	Documents   *api.DocumentsApiService
	QueryApi    *api.QueriesApiService
	Users       *api.UsersApiService
}

// NewClient creates a new Rockset client. Either the FromEnv() or WithAPIKey() option has to be supplied.
func NewClient(options ...RockOption) (*RockClient, error) {
	rc := &RockClient{
		apiServer:  DefaultAPIServer,
		httpClient: http.DefaultClient,
	}

	for _, o := range options {
		o(rc)
	}

	if err := rc.Validate(); err != nil {
		return nil, err
	}

	cfg := api.NewConfiguration()
	cfg.BasePath = rc.apiServer
	cfg.Version = Version
	cfg.HTTPClient = rc.httpClient
	rc.common.Client = api.ApiClient(cfg, rc.apiKey)

	rc.ApiKeys = (*api.ApiKeysApiService)(&rc.common)
	rc.Collection = (*api.CollectionsApiService)(&rc.common)
	rc.Integration = (*api.IntegrationsApiService)(&rc.common)
	rc.Documents = (*api.DocumentsApiService)(&rc.common)
	rc.QueryApi = (*api.QueriesApiService)(&rc.common)
	rc.Users = (*api.UsersApiService)(&rc.common)

	return rc, nil
}

// Query executes a query request against Rockset
func (rc *RockClient) Query(request api.QueryRequest) (api.QueryResponse, *http.Response, error) {
	return rc.QueryApi.Query(request)
}

// Validate validates and sets the Rockset client configuration options
func (rc *RockClient) Validate() error {
	if rc.apiKey == "" {
		return fmt.Errorf("an API key must be specified")
	}
	if rc.apiServer == "" {
		return fmt.Errorf("an API server must be specified")
	}

	u, err := url.Parse(rc.apiServer)
	if err != nil {
		return fmt.Errorf("failed to parse API server %s: %w", rc.apiServer, err)
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	rc.apiServer = u.String()

	if rc.timeout != 0 {
		rc.httpClient.Timeout = rc.timeout
	}

	return nil
}

// RockOption is the type for RockClient options
type RockOption func(rc *RockClient)

// FromEnv sets API key and API server from the environment variables ROCKSET_APIKEY and ROCKSET_APISERVER
func FromEnv() RockOption {
	return func(rc *RockClient) {
		rc.apiKey = os.Getenv("ROCKSET_APIKEY")

		if server := os.Getenv("ROCKSET_APISERVER"); server != "" {
			rc.apiServer = server
		}
	}
}

// WithAPIKey sets the API key to key
func WithAPIKey(key string) RockOption {
	return func(rc *RockClient) {
		rc.apiKey = key
	}
}

// WithAPIServer sets the API server
func WithAPIServer(s string) RockOption {
	return func(rc *RockClient) {
		rc.apiServer = s
	}
}

// WithHTTPClient sets the HTTP client. Without this option RockClient uses the http.DefaultClient,
// which does not have a timeout.
func WithHTTPClient(c *http.Client) RockOption {
	return func(rc *RockClient) {
		rc.httpClient = c
	}
}

// WithTimeout sets the HTTP client timeout, and will override any value set using using the WithHTTPClient() option
func WithTimeout(t time.Duration) RockOption {
	return func(rc *RockClient) {
		rc.timeout = t
	}
}

// Create a Client object to securely connect to Rockset using an API key
// Optionally, an alternate API server host can also be provided.
//
// Deprecated: this function has been superseded by NewClient()
func Client(apiKey string, apiServer string) *RockClient {
	// TODO read from credentials file if it exists
	if apiKey == "" {
		log.Fatal("apiKey needs to be specified")
	}

	if apiServer == "" {
		apiServer = "https://api.rs2.usw2.rockset.com"
	}

	if !strings.HasPrefix(apiServer, "http://") && !strings.HasPrefix(apiServer, "https://") {
		apiServer = "https://" + apiServer
	}

	c := &RockClient{}
	cfg := api.NewConfiguration()
	cfg.BasePath = apiServer
	cfg.Version = Version
	c.common.Client = api.ApiClient(cfg, apiKey)

	// API Services
	c.ApiKeys = (*api.ApiKeysApiService)(&c.common)
	c.Collection = (*api.CollectionsApiService)(&c.common)
	c.Integration = (*api.IntegrationsApiService)(&c.common)
	c.Documents = (*api.DocumentsApiService)(&c.common)
	c.QueryApi = (*api.QueriesApiService)(&c.common)
	c.Users = (*api.UsersApiService)(&c.common)

	return c
}
