package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

func TestRockClient_CreateQueryLambda(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	ws := "acc"
	name := randomName("ql")

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	ql, err := rc.CreateQueryLambda(ctx, ws, name, "SELECT 1",
		option.WithDefaultParameter("", "", ""), option.WithQueryLambdaDescription(description()))
	require.NoError(t, err)

	defer func() {
		err := rc.DeleteQueryLambda(ctx, ws, name)
		assert.NoError(t, err)
	}()

	assert.Equal(t, name, ql.GetName())

	ql, err = rc.UpdateQueryLambda(ctx, ws, name, "SELECT 2",
		option.WithDefaultParameter("dummy", "string", "foo"))
	assert.NoError(t, err)
}

// test fixtures
const (
	qlName    = "events"
	qlVersion = "1921bcad817b5cc2"
	qlTag     = "test"
)

func TestRockClient_GetQueryLambdaVersionByTag(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	version, err := rc.GetQueryLambdaVersionByTag(ctx, persistentWorkspace, qlName, qlTag)
	require.NoError(t, err)
	assert.Equal(t, qlVersion, version.Version.GetVersion())
}

func TestRockClient_GetQueryLambdaVersion(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	version, err := rc.GetQueryLambdaVersion(ctx, persistentWorkspace, qlName, qlVersion)
	require.NoError(t, err)
	assert.Equal(t, qlName, version.GetName())
}

func TestRockClient_ListQueryLambdas(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	rc, err := rockset.NewClient()
	require.NoError(t, err)

	lambdas, err := rc.ListQueryLambdas(ctx)
	require.NoError(t, err)

	for _, l := range lambdas {
		t.Logf("lambda: %s", l.GetName())
	}
}

func TestRockClient_ListQueryLambdas_workspace(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	lambdas, err := rc.ListQueryLambdas(ctx, option.WithQueryLambdaWorkspace(persistentWorkspace))
	require.NoError(t, err)

	for _, l := range lambdas {
		assert.Equal(t, persistentWorkspace, l.GetWorkspace())
	}
}

func TestRockClient_ListQueryLambdaVersions(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	versions, err := rc.ListQueryLambdaVersions(ctx, persistentWorkspace, qlName)
	require.NoError(t, err)

	for _, l := range versions {
		t.Logf("version: %s", l.GetVersion())
	}
}

func TestRockClient_ListQueryLambdaTags(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	tags, err := rc.ListQueryLambdaTags(ctx, persistentWorkspace, qlName)
	require.NoError(t, err)

	for _, tag := range tags {
		t.Logf("tag: %s", tag.GetTagName())
	}
}
