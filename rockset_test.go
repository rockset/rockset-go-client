package rockset_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
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

func testClient(t *testing.T) *rockset.RockClient {
	rc, err := rockset.NewClient(rockset.WithUserAgent("rockset-go-integration-tests"))
	require.NoError(t, err)
	return rc
}

func debugClient(t *testing.T) *rockset.RockClient {
	rc, err := rockset.NewClient(rockset.WithUserAgent("rockset-go-integration-tests"), rockset.WithHTTPDebug())
	require.NoError(t, err)
	return rc
}

// these are used for testing when a persistent value is needed
const buildNum = "CIRCLE_BUILD_NUM"
const persistentWorkspace = "persistent"
const persistentCollection = "snp"
const persistentAlias = "getalias"

func description() string {
	num, found := os.LookupEnv(buildNum)
	if !found {
		num = "dev"
	}
	return fmt.Sprintf("created by terraform integration test run %s", num)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset creates a random string with length and charset
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// String creates a random string with length
func randomString(length int) string {
	return stringWithCharset(length, charset)
}

func randomName(prefix string) string {
	num, found := os.LookupEnv(buildNum)
	if !found {
		if user, found := os.LookupEnv("USER"); found {
			num = strings.ToLower(user)
		} else {
			num = "dev"
		}
	}

	return fmt.Sprintf("go_%s_%s_%s", num, prefix, randomString(6))
}

// TestTemplate is used as a copypasta for new tests
func TestTemplate(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc := testClient(t)

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

const USW2A1 = "api.usw2a1.rockset.com"

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

	rc, err := rockset.NewClient(rockset.WithAPIServer(fmt.Sprintf("https://%s/", USW2A1)))
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
	err := os.Setenv(rockset.APIServerEnvironmentVariableName, USW2A1)
	require.NoError(t, err)

	rc := testClient(t)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	log.Debug().Str("org", org.GetDisplayName())
}

func TestRockClient_Ping(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc := testClient(t)

	err := rc.Ping(ctx)
	require.NoError(t, err)
}
