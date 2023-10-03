package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client"
)

// SkipUnlessIntegrationTest is helper function to skip unless ROCKSET_APIKEY is set
func SkipUnlessIntegrationTest(t *testing.T) {
	_ = SkipUnlessEnvSet(t, rockset.APIKeyEnvironmentVariableName)
}

func SkipUnlessEnvSet(t *testing.T, envName string) string {
	env := os.Getenv(envName)
	if env == "" {
		t.Skipf("skipping as %s is not set", envName)
	}

	return env
}

func SkipUnlessDocker(t *testing.T) {
	_, err := os.Stat("/var/run/docker.sock")
	if os.IsNotExist(err) {
		t.Skip("docker socket not present, skipping")
	}
	assert.NoError(t, err)
}
