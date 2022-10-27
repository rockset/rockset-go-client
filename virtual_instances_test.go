package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRockClient_ListVirtualInstances(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	rc := testClient(t)

	vis, err := rc.ListVirtualInstances(ctx)
	require.NoError(t, err)
	assert.NotEmpty(t, vis)

	for _, vi := range vis {
		t.Logf("vi %s: %s %s", vi.GetId(), vi.GetState(), vi.GetCurrentSize())
		assert.Equal(t, "SMALL", vi.GetCurrentSize())
		assert.Equal(t, "SMALL", vi.GetDesiredSize())
	}
}
