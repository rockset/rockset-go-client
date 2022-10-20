package rockset_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
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

func skipUnlessDocker(t *testing.T) {
	_, err := os.Stat("/var/run/docker.sock")
	if os.IsNotExist(err) {
		t.Skip("docker socket not present, skipping")
	}
	assert.NoError(t, err)
}

// helper function to create a context with a zerolog logger
func testCtx() context.Context {
	return testCtxWithLevel(zerolog.WarnLevel)
}

func testCtxWithLevel(lvl zerolog.Level) context.Context {
	ctx := context.Background()
	console := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	log := zerolog.New(console).Level(lvl).With().Timestamp().Logger()

	return log.WithContext(ctx)
}

// these are used for testing when a persistent value is needed
const buildNum = "CIRCLE_BUILD_NUM"
const persistentWorkspace = "persistent"
const persistentCollection = "snp"

func description() string {
	num, found := os.LookupEnv(buildNum)
	if !found {
		num = "dev"
	}
	return fmt.Sprintf("created by terraform integration test run %s", num)
}

func randomWorkspace(t *testing.T) string {
	n, found := os.LookupEnv("CIRCLE_BUILD_NUM")
	if !found {
		return "go_dev"
	}
	num, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		t.Errorf("failed to convert %s to int: %v", n, err)
	}

	return fmt.Sprintf("go_%d", num)
}

func randomName(t *testing.T, prefix string) string {
	n, found := os.LookupEnv(buildNum)
	if !found {
		return fmt.Sprintf("go_%s_dev", prefix)
	}
	num, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		t.Errorf("failed to convert %s to int: %v", n, err)
	}

	return fmt.Sprintf("go_%s_%d", prefix, num)
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
