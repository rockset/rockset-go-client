package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"github.com/rockset/rockset-go-client/internal/test"
)

func TestRockClient_GetOrganization(t *testing.T) {
	ctx := test.Context()
	rc, _ := vcrTestClient(t, t.Name())

	org, err := rc.GetOrganization(ctx)
	require.NoError(t, err)

	assert.Equal(t, "rockset-circleci", org.GetId())
}
