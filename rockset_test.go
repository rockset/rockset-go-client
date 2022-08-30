package rockset_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

// APIKeyEnvironmentVariableName is the environment variable name for the API key
const APIKeyEnvironmentVariableName = "ROCKSET_APIKEY" //nolint

// APIServerEnvironmentVariableName is the environment variable name for the API server
const APIServerEnvironmentVariableName = "ROCKSET_APISERVER"

// helper function to skip unless ROCKSET_APIKEY is set
func skipUnlessIntegrationTest(t *testing.T) {
	_ = skipUnlessEnvSet(t, rockset.APIKeyEnvironmentVariableName)
}

func skipUnlessEnvSet(t *testing.T, envName string) string {
	env := os.Getenv(envName)
	if env == "" {
		t.Skipf("skipping as %s is not set", envName)
	}

	return env
}

// helper function to create a context with a zerolog logger
func testCtx() context.Context {
	ctx := context.Background()
	console := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	log := zerolog.New(console).Level(zerolog.TraceLevel).With().Timestamp().Logger()

	return log.WithContext(ctx)
}

// TestTemplate is used as a copypasta for new tests
func TestTemplate(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	log.Debug().Str("org", org.GetDisplayName())
}

func TestRockClient_withAPIKey(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	key := os.Getenv(rockset.APIKeyEnvironmentVariableName)
	err := os.Unsetenv(rockset.APIKeyEnvironmentVariableName)
	require.NoError(t, err)

	defer func() {
		if err = os.Setenv(rockset.APIKeyEnvironmentVariableName, key); err != nil {
			panic("failed to reset environment")
		}
	}()

	rc, err := rockset.NewClient(rockset.WithAPIKey(key))
	require.NoError(t, err)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	log.Debug().Str("org", org.GetDisplayName())
}

func TestRockClient_withoutAPIServer(t *testing.T) {
	// this is messing with the environment
	oldEnv, set := os.LookupEnv(rockset.APIServerEnvironmentVariableName)
	if set {
		defer func() {
			if err := os.Setenv(rockset.APIServerEnvironmentVariableName, oldEnv); err != nil {
				t.Errorf("failed to reset environment variable %s: %v", rockset.APIServerEnvironmentVariableName, err)
			}
		}()
		err := os.Unsetenv(rockset.APIServerEnvironmentVariableName)
		require.NoError(t, err)
	}

	_, err := rockset.NewClient()
	if assert.Error(t, err) {
		assert.Equal(t, "you must specify an API server", err.Error())
	}
}

func TestRockClient_withAPIServer(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	// this is messing with the environment
	oldEnv := os.Getenv(rockset.APIServerEnvironmentVariableName)
	defer func() {
		if err := os.Setenv(rockset.APIServerEnvironmentVariableName, oldEnv); err != nil {
			t.Errorf("failed to reset environment variable %s: %v", rockset.APIServerEnvironmentVariableName, err)
		}
	}()
	err := os.Unsetenv(rockset.APIServerEnvironmentVariableName)
	require.NoError(t, err)

	rc, err := rockset.NewClient(rockset.WithAPIServer("https://api.use1a1.rockset.com/"))
	require.NoError(t, err)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	log.Debug().Str("org", org.GetDisplayName())
}

func TestRockClient_withAPIServerEnv(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	// this is messing with the environment
	oldEnv := os.Getenv(rockset.APIServerEnvironmentVariableName)
	defer func() {
		if err := os.Setenv(rockset.APIServerEnvironmentVariableName, oldEnv); err != nil {
			t.Logf("failed to reset environment variable %s: %v", rockset.APIServerEnvironmentVariableName, err)
		}
	}()
	err := os.Setenv(rockset.APIServerEnvironmentVariableName, "api.use1a1.rockset.com")
	require.NoError(t, err)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	log.Debug().Str("org", org.GetDisplayName())
}

func TestRockClient_Ping(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	err = rc.Ping(ctx)
	require.NoError(t, err)
}
