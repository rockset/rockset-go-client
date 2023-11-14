package rockset

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/retry"
	"github.com/rockset/rockset-go-client/wait"
)

const (
	// APIKeyEnvironmentVariableName is the environment variable name for the API key
	APIKeyEnvironmentVariableName = "ROCKSET_APIKEY" //nolint
	// APIServerEnvironmentVariableName is the environment variable name for the API server
	APIServerEnvironmentVariableName = "ROCKSET_APISERVER"
	// DefaultUserAgent is the user agent string the Rockset go client will use for REST API calls.
	DefaultUserAgent = "rockset-go-client"
	// HeaderVersionName is the name of the HTTP header the go client sets which contains the client version used.
	HeaderVersionName = "x-rockset-version"
)

// RockConfig contains the configurable options for the RockClient.
type RockConfig struct {
	// Retrier is the retry function used to retry API calls.
	retry.Retrier
	// Wait is the helper to wait for resources
	Wait *wait.Waiter
	// APIKey is the API key to use for authentication
	APIKey string
	// APIServer is the API server to connect to
	APIServer string
	// Token is an OAuth2 ID token
	Token string
	// Organization is the Rockset org, which is required when authenticating using Token
	Organization string
	cfg          *openapi.Configuration
	APIClient    *openapi.APIClient
}

// RockClient is the client struct for making APi calls to Rockset.
type RockClient struct {
	*openapi.APIClient
	RockConfig
}

// NewClient creates a new Rockset client.
//
// Accessing the online database requires an API key and an API server to connect to,
// which can be provided by through the ROCKSET_APIKEY and ROCKSET_APISERVER environment variables,
//
//	c, err := rockset.NewClient()
//
// or they can be explicitly set using the WithAPIKey() and WithAPIServer() options.
//
//	c, err := rockset.NewClient(rockset.WithAPIKey("..."), rockset.WithAPIServer("..."))
func NewClient(options ...RockOption) (*RockClient, error) {
	cfg := openapi.NewConfiguration()
	cfg.UserAgent = DefaultUserAgent
	cfg.AddDefaultHeader(HeaderVersionName, Version)
	// TODO should the default http client be tuned?
	cfg.HTTPClient = &http.Client{}

	rc := RockConfig{
		cfg:       cfg,
		Retrier:   retry.NewExponential(),
		APIKey:    os.Getenv(APIKeyEnvironmentVariableName),
		APIServer: os.Getenv(APIServerEnvironmentVariableName),
	}

	for _, o := range options {
		o(&rc)
	}

	// if we have an APIClient then we assume it is a fake, and no further setup is required
	if rc.APIClient != nil {
		client := RockClient{
			RockConfig: rc,
			APIClient:  rc.APIClient,
		}

		client.Wait = wait.New(&client)
		return &client, nil
	}

	if rc.APIServer == "" {
		return nil, ErrNoAPIServer
	}

	u, err := url.Parse(rc.APIServer)
	if err != nil {
		return nil, err
	}

	if u.Host != "" { // https:// was used
		rc.APIServer = u.Host
	} else if u.String() != "" { // plain hostname was used
		rc.APIServer = u.String()
	} else {
		return nil, fmt.Errorf("%s is not a valid URL", rc.APIServer)
	}

	cfg.Host = rc.APIServer
	cfg.Scheme = "https" // we do not allow setting the scheme from the URL as we only support HTTPS

	if rc.APIKey == "" && rc.Token == "" {
		return nil, ErrNoAPICredentials
	} else if rc.APIKey != "" && rc.Token != "" {
		return nil, ErrDuplicateCredentials
	} else if rc.APIKey != "" {
		cfg.AddDefaultHeader("Authorization", "apikey "+rc.APIKey)
	} else {
		cfg.AddDefaultHeader("Authorization", "bearer "+rc.Token)
		cfg.AddDefaultHeader("organization", rc.Organization)
	}

	client := RockClient{
		RockConfig: rc,
		APIClient:  openapi.NewAPIClient(rc.cfg),
	}
	client.Wait = wait.New(&client)

	return &client, nil
}

var (
	ErrNoAPICredentials     = errors.New("no API credentials provided")
	ErrDuplicateCredentials = errors.New("duplicate API credentials provided")
	ErrNoAPIServer          = errors.New("no API server provided")
)

// RockOption is the type for RockClient options.
type RockOption func(rc *RockConfig)

// WithAPIKey sets the API key to use, and removes the bearer token.
func WithAPIKey(apiKey string) RockOption {
	return func(rc *RockConfig) {
		rc.APIKey = apiKey
		rc.Token = ""
		rc.Organization = ""
	}
}

// WithAPIServer sets the API server to connect to, and override the ROCKSET_APISERVER.
func WithAPIServer(server string) RockOption {
	return func(rc *RockConfig) {
		rc.APIServer = server
	}
}

// WithAPIClient sets the openapi.APIClient to use, which can be used in testing to use a fake client.
//
//	f := fake.NewAPIClient()
//	rc, err := rockset.NewRockClient(rockset.WithAPIClient(f)
func WithAPIClient(c *openapi.APIClient) RockOption {
	return func(rc *RockConfig) {
		rc.APIClient = c
	}
}

// WithBearerToken uses authorization using a bearer token, which requires specifying the organization.
// Removes the API key, if set.
func WithBearerToken(token, org string) RockOption {
	return func(rc *RockConfig) {
		rc.Token = token
		rc.Organization = org
		rc.APIKey = ""
	}
}

// WithHTTPClient sets the HTTP client. Without this option RockClient uses the http.DefaultClient.
func WithHTTPClient(c *http.Client) RockOption {
	return func(rc *RockConfig) {
		rc.cfg.HTTPClient = c
	}
}

// WithCustomHeader sets a custom HTTP header for all Rockset API calls.
func WithCustomHeader(key, value string) RockOption {
	return func(rc *RockConfig) {
		rc.cfg.AddDefaultHeader(key, value)
	}
}

// WithRetry sets the Retrier the RockClient uses to retry requests which return an Error that can be retried.
// The errors which can be retried are configurable using the Exponential field RetryableErrorCheck.
//
//	exp := rockset.Exponential{
//		RetryableErrorCheck: func(err error) bool {
//			return error.Is(err, io.ErrUnexpectedEOF)
//		}
//	}
//	rc, err := rockset.NewClient(rockset.WithRetry(exp))
//	// handle error
//	err = rc.Retry(ctx, func() error{
//		// call that will be retried if it returns io.ErrUnexpectedEOF
//	})
//
// This would retry all io.ErrUnexpectedEOF errors
func WithRetry(r retry.Retrier) RockOption {
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

// WithUserAgent sets the user agent string. Used by the Rockset terraform provider.
// If not set, it defaults to DefaultUserAgent.
func WithUserAgent(name string) RockOption {
	return func(rc *RockConfig) {
		rc.cfg.UserAgent = name
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

// Ping connects to the Rockset API server and can be used to verify connectivity. It does not
// require authentication, so to test that use the GetOrganization() method instead.
func (rc *RockClient) Ping(ctx context.Context) error {
	c := rc.RockConfig.cfg.HTTPClient
	u := url.URL{
		Scheme: rc.RockConfig.cfg.Scheme,
		Host:   rc.RockConfig.cfg.Host,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	// needs special handling as the VCR recorder inspects this field
	req.Header.Set(HeaderVersionName, Version)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected return code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var ping struct {
		Message string `json:"message"`
	}
	err = json.Unmarshal(body, &ping)
	if err != nil {
		return err
	}

	log := zerolog.Ctx(ctx)
	log.Debug().Msgf("ping response: %s", ping.Message)

	const goRockset = "GO ROCKSET!"
	if ping.Message != goRockset {
		return fmt.Errorf("unexpected message: %s", ping.Message)
	}

	return nil
}

const (
	entityRegexpString = "^[[:alnum:]][[:alnum:]-_]*$"
	maxEntityLength    = 100
)

var (
	entityRegexp = regexp.MustCompile(entityRegexpString)
)

// ValidEntityName is used to validate if a name is valid for a workspace, collection, etc
func ValidEntityName(name string) error {
	if len(name) > maxEntityLength {
		return fmt.Errorf("name must be less than %d characters", maxEntityLength)
	}

	if !entityRegexp.MatchString(name) {
		return fmt.Errorf("invalid name, must match `%s`", entityRegexpString)
	}

	return nil
}
