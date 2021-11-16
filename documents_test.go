package rockset_test

import (
	"testing"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func TestRockClient_PatchDocuments(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	type doc struct {
		Foo string `json:"foo"`
		Bar bool   `json:"bar"`
	}

	var docs = []interface{}{
		structs.Map(doc{Foo: "foo"}),
	}

	res, err := rc.AddDocuments(ctx, "tests", "patch", docs)
	require.NoError(t, err)
	require.Len(t, res, 1)

	patches := []rockset.PatchDocument{
		{
			ID: *res[0].Id,
			Patches: []rockset.PatchOperation{
				{
					Op:    "replace",
					Path:  "/Bar",
					Value: true,
				},
			},
		},
	}
	res, err = rc.PatchDocuments(ctx, "tests", "patch", patches)
	require.NoError(t, err)
	require.Len(t, res, 1)
	assert.Equal(t, "PATCHED", *res[0].Status)
	// TODO: verify that it is actually patched too?
}
