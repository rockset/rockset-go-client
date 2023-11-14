package rockset_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/fake"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/openapi"
)

// TestTemplate is used as a copypasta for new tests
func TestTemplate(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()
	log := zerolog.Ctx(ctx)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	log.Debug().Str("org", org.GetDisplayName())
}

// TestFake is an example on how to use a fake RockClient to test. It takes a lot of setup,
// but gives you more control than the VCR client.
func TestFake(t *testing.T) {
	f := fake.NewAPIClient()
	org := f.OrganizationsApi.(*fake.FakeOrganizationsApi)
	org.GetOrganizationReturns(openapi.ApiGetOrganizationRequest{
		ApiService: org,
	})
	org.GetOrganizationExecuteReturns(&openapi.OrganizationResponse{
		Data: &openapi.Organization{
			DisplayName: openapi.PtrString("rockset"),
		},
	}, &http.Response{}, nil)

	rc, err := rockset.NewClient(rockset.WithAPIClient(f))
	assert.NoError(t, err)

	o, err := rc.GetOrganization(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "rockset", o.GetDisplayName())
}

type RockOptionSuite struct {
	suite.Suite
	savedAPIKey    string
	savedAPIServer string
}

func TestRockOptions(t *testing.T) {
	s := RockOptionSuite{
		// save API key & server
		savedAPIKey:    os.Getenv(rockset.APIKeyEnvironmentVariableName),
		savedAPIServer: os.Getenv(rockset.APIServerEnvironmentVariableName),
	}
	suite.Run(t, &s)
}

func (s *RockOptionSuite) TearDownSuite() {
	// restore API key & server
	if s.savedAPIKey != "" {
		s.NoError(os.Setenv(rockset.APIKeyEnvironmentVariableName, s.savedAPIKey))
	}
	if s.savedAPIServer != "" {
		s.NoError(os.Setenv(rockset.APIServerEnvironmentVariableName, s.savedAPIServer))
	}
}

func (s *RockOptionSuite) SetupTest() {
	// clear env before each test
	s.NoError(os.Unsetenv(rockset.APIKeyEnvironmentVariableName))
	s.NoError(os.Unsetenv(rockset.APIServerEnvironmentVariableName))
}

func (s *RockOptionSuite) TestMissingCreds() {
	_, err := rockset.NewClient(rockset.WithAPIServer("server"))
	s.ErrorIs(err, rockset.ErrNoAPICredentials)
}

func (s *RockOptionSuite) TestMissingServer() {
	_, err := rockset.NewClient(rockset.WithAPIKey("key"))
	s.ErrorIs(err, rockset.ErrNoAPIServer)
}

func (s *RockOptionSuite) TestLastCredWins() {
	rc, err := rockset.NewClient(
		rockset.WithAPIServer("server"),
		rockset.WithBearerToken("token", "org"),
		rockset.WithAPIKey("key"),
	)
	s.NoError(err)
	s.Equal("key", rc.APIKey)
	s.Equal("", rc.Token)
	s.Equal("", rc.Organization)
}

func (s *RockOptionSuite) TestExplicitKeyOverridesEnv() {
	s.NoError(os.Setenv(rockset.APIKeyEnvironmentVariableName, "env"))
	rc, err := rockset.NewClient(rockset.WithAPIServer("server"), rockset.WithAPIKey("key"))
	s.NoError(err)
	s.Equal("key", rc.APIKey)
}

func (s *RockOptionSuite) TestExplicitServerOverridesEnv() {
	s.NoError(os.Setenv(rockset.APIKeyEnvironmentVariableName, "env"))
	rc, err := rockset.NewClient(rockset.WithAPIServer("server"), rockset.WithAPIKey("key"))
	s.NoError(err)
	s.Equal("server", rc.APIServer)
}

func (s *RockOptionSuite) TestBearerToken() {
	token := "bearer token"
	org := "rockset org"

	rc, err := rockset.NewClient(rockset.WithBearerToken(token, org), rockset.WithAPIServer("server"))
	s.NoError(err)
	s.Equal(token, rc.Token)
	s.Equal(org, rc.Organization)
	s.Equal("", rc.APIKey)
}

func (s *RockOptionSuite) TestBearerTokenWithEnv() {
	s.NoError(os.Setenv(rockset.APIKeyEnvironmentVariableName, "env"))
	token := "bearer token"
	org := "rockset org"

	rc, err := rockset.NewClient(rockset.WithBearerToken(token, org), rockset.WithAPIServer("server"))
	s.NoError(err)
	s.Equal(token, rc.Token)
	s.Equal(org, rc.Organization)
	s.Equal("", rc.APIKey)
}

func TestRockClient_Ping(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	err := rc.Ping(ctx)
	require.NoError(t, err)
}

func TestValidCollectionName(t *testing.T) {
	tests := []struct {
		name   string
		errStr string
	}{
		{"_123", "invalid name, must match `^[[:alnum:]][[:alnum:]-_]*$`"},
		{"abc", ""},
		{
			"a1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"name must be less than 100 characters",
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			err := rockset.ValidEntityName(tst.name)
			if tst.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tst.errStr, err.Error())
			}
		})
	}
}
