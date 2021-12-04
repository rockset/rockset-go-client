package rockset_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client"
)

// helper function to skip unless ROCKSET_APIKEY is set
func skipUnlessIntegrationTest(t *testing.T) {
	if os.Getenv(rockset.APIKeyEnvironmentVariableName) == "" {
		t.Skipf("skipping as %s is not set", rockset.APIKeyEnvironmentVariableName)
	}
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

func TestWithAPIKey(t *testing.T) {
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

const use1a1 = "https://api.use1a1.rockset.com/"

func TestWithAPIServer(t *testing.T) {
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
	err := os.Unsetenv(rockset.APIServerEnvironmentVariableName)
	require.NoError(t, err)

	rc, err := rockset.NewClient(rockset.WithAPIServer(use1a1), rockset.WithHTTPDebug())
	require.NoError(t, err)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	log.Debug().Str("org", org.GetDisplayName())
}

func TestRockClientWithAPIServerEnv(t *testing.T) {
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
	err := os.Setenv(rockset.APIServerEnvironmentVariableName, use1a1)
	require.NoError(t, err)

	rc, err := rockset.NewClient(rockset.WithHTTPDebug())
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
