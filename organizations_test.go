package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRockClient_GetOrganization(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc := testClient(t)

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	assert.Equal(t, "rockset-circleci", org.GetId())
}
