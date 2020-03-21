package rockset

import (
	"encoding/json"
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
const Version = "0.9.1"

// DefaultAPIServer is the default Rockset API server to use
const DefaultAPIServer = "https://api.rs2.usw2.rockset.com"

const APIKeyEnvironmentVariableName = "ROCKSET_APIKEY"
const APIServerEnvironmentVariableName = "ROCKSET_APISERVER"

type RockClient struct {
	apiServer  string
	apiKey     string
	timeout    time.Duration
	httpClient *http.Client

	common api.Service

	// API Services
	ApiKeys       *api.ApiKeysApiService
	Collection    *api.CollectionsApiService
	Integration   *api.IntegrationsApiService
	Documents     *api.DocumentsApiService
	QueryApi      *api.QueriesApiService
	Users         *api.UsersApiService
	Organizations *api.OrganizationsApiService
	QueryLambdas  *api.QueryLambdasApiService
	Workspaces    *api.WorkspacesApiService
}

// NewClient creates a new Rockset client.
//
// Accessing the online database requires an API key, which you either have to supply through
// the ROCKSET_APIKEY environment variable and pass the FromEnv() option
//	c, err := rockset.NewClient(rockset.FromEnv())
// or explicitly using the WithAPIKey() option
//	c, err := rockset.NewClient(rockset.WithAPIKey("..."))
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
	rc.Documents = (*api.DocumentsApiService)(&rc.common)
	rc.Integration = (*api.IntegrationsApiService)(&rc.common)
	rc.QueryApi = (*api.QueriesApiService)(&rc.common)
	rc.Users = (*api.UsersApiService)(&rc.common)
	rc.Organizations = (*api.OrganizationsApiService)(&rc.common)
	rc.QueryLambdas = (*api.QueryLambdasApiService)(&rc.common)
	rc.Workspaces = (*api.WorkspacesApiService)(&rc.common)

	return rc, nil
}

// Query executes a query request against Rockset
func (rc *RockClient) Query(request api.QueryRequest) (api.QueryResponse, *http.Response, error) {
	return rc.QueryApi.Query(request)
}

// Organization returns the organization the RockClient belongs to
func (rc *RockClient) Organization() (api.Organization, *http.Response, error) {
	org, resp, err := rc.Organizations.Get()
	if err != nil {
		return api.Organization{}, resp, err
	}
	return *org.Data, resp, err
}

// GetWorkspace gets the workspace with name
func (rc *RockClient) GetWorkspace(name string) (api.Workspace, *http.Response, error) {
	w, resp, err := rc.Workspaces.Get(name)
	if err != nil {
		return api.Workspace{}, resp, err
	}
	return *w.Data, resp, err
}

// ListWorkspaces list all workspaces
func (rc *RockClient) ListWorkspaces() ([]api.Workspace, *http.Response, error) {
	w, resp, err := rc.Workspaces.List()
	if err != nil {
		return nil, resp, err
	}
	return w.Data, resp, err
}

// CreateWorkspace creates a new workspace
func (rc *RockClient) CreateWorkspace(name, description string) (api.Workspace, *http.Response, error) {
	w, resp, err := rc.Workspaces.Create(api.CreateWorkspaceRequest{Name: name, Description: description})
	if err != nil {
		return api.Workspace{}, resp, err
	}
	return *w.Data, resp, err
}

// DeleteWorkspace deletes the workspace with name
func (rc *RockClient) DeleteWorkspace(name string) (api.Workspace, *http.Response, error) {
	w, resp, err := rc.Workspaces.Delete(name)
	if err != nil {
		return api.Workspace{}, resp, err
	}
	return *w.Data, resp, err
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

// FromEnv sets API key and API server from the environment variables ROCKSET_APIKEY and ROCKSET_APISERVER,
// and if ROCKSET_APISERVER is not set, it will use the default API server.
func FromEnv() RockOption {
	return func(rc *RockClient) {
		rc.apiKey = os.Getenv(APIKeyEnvironmentVariableName)

		if server := os.Getenv(APIServerEnvironmentVariableName); server != "" {
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

// AsRocksetError takes an error returned from an API call and returns the underlying error message
func AsRocksetError(err error) (api.ErrorModel, bool) {
	var em api.ErrorModel
	if e, ok := err.(api.GenericSwaggerError); ok {
		if err = json.Unmarshal(e.Body(), &em); err == nil {
			return em, true
		}
	}
	return em, false
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
