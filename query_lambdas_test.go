package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

func TestRockClient_CreateQueryLambda(t *testing.T) {
	rc, randomName := vcrTestClient(t, t.Name())

	ctx := test.Context()
	ws := "acc"
	name := randomName("ql")

	ql, err := rc.CreateQueryLambda(ctx, ws, name, "SELECT 1",
		option.WithDefaultParameter("", "", ""), option.WithQueryLambdaDescription(test.Description()))
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
	qlVersion = "0eb7783c81ef339e"
	qlTag     = "test"
)

func TestRockClient_GetQueryLambdaVersionByTag(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	version, err := rc.GetQueryLambdaVersionByTag(ctx, test.PersistentWorkspace, qlName, qlTag)
	require.NoError(t, err)
	assert.Equal(t, qlVersion, version.Version.GetVersion())
}

func TestRockClient_GetQueryLambdaVersion(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	version, err := rc.GetQueryLambdaVersion(ctx, test.PersistentWorkspace, qlName, qlVersion)
	require.NoError(t, err)
	assert.Equal(t, qlName, version.GetName())
}

func TestRockClient_ListQueryLambdas(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	lambdas, err := rc.ListQueryLambdas(ctx)
	require.NoError(t, err)

	for _, l := range lambdas {
		t.Logf("lambda: %s", l.GetName())
	}
}

func TestRockClient_ListQueryLambdas_workspace(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	lambdas, err := rc.ListQueryLambdas(ctx, option.WithQueryLambdaWorkspace(test.PersistentWorkspace))
	require.NoError(t, err)

	for _, l := range lambdas {
		assert.Equal(t, test.PersistentWorkspace, l.GetWorkspace())
	}
}

func TestRockClient_ListQueryLambdaVersions(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	versions, err := rc.ListQueryLambdaVersions(ctx, test.PersistentWorkspace, qlName)
	require.NoError(t, err)

	for _, l := range versions {
		t.Logf("version: %s", l.GetVersion())
	}
}

func TestRockClient_ListQueryLambdaTags(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	tags, err := rc.ListQueryLambdaTags(ctx, test.PersistentWorkspace, qlName)
	require.NoError(t, err)

	for _, tag := range tags {
		t.Logf("tag: %s", tag.GetTagName())
	}
}
