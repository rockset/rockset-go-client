package rockset_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
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

func TestRockClient_withAPIKey(t *testing.T) {
	test.SkipUnlessIntegrationTest(t)

	ctx := test.Context()
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
	test.SkipUnlessIntegrationTest(t)
	// TODO this should use VCR too

	ctx := test.Context()
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
	rc, _ := vcrTestClient(t, t.Name())

	ctx := test.Context()
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

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	log.Debug().Str("org", org.GetDisplayName())
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
