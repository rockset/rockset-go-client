package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRockClient_GetOrganization(t *testing.T) {
	ctx := testCtx()
	rc, _ := vcrClient(t, t.Name())

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	assert.Equal(t, "rockset-circleci", org.GetId())
}
