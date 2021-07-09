package rockset_test

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

func TestRockClient_CreateQueryLambda(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	ql, err := rc.CreateQueryLambda(ctx, "commons", "qlTest", "SELECT 1",
		option.WithDefaultParameter("", "", ""))
	require.NoError(t, err)

	defer func() {
		err := rc.DeleteQueryLambda(ctx, "commons", "qlTest")
		assert.NoError(t, err)
	}()

	assert.Equal(t, "qlTest", *ql.Name)

	ql, err = rc.UpdateQueryLambda(ctx, "commons", "qlTest", "SELECT 2",
		option.WithDefaultParameter("dummy", "string", "foo"))
	assert.NoError(t, err)
}

func TestRockClient_GetQueryLambdaVersionByTag(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	version, err := rc.GetQueryLambdaVersionByTag(ctx, "commons", "events", "v2")
	require.NoError(t, err)
	assert.Equal(t, "2", *version.Version.Version)
}

func TestRockClient_GetQueryLambdaVersion(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	version, err := rc.GetQueryLambdaVersion(ctx, "commons", "events", "2")
	require.NoError(t, err)
	assert.Equal(t, "events", *version.Name)
}

func TestRockClient_ListQueryLambdas(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	lambdas, err := rc.ListQueryLambdas(ctx)
	require.NoError(t, err)

	for _, l := range lambdas {
		log.Printf("lambda: %s", *l.Name)
	}
}

func TestRockClient_ListQueryLambdas_workspace(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	lambdas, err := rc.ListQueryLambdas(ctx, option.WithQueryLambdaWorkspace("commons"))
	require.NoError(t, err)

	for _, l := range lambdas {
		log.Printf("lambda: %s", *l.Name)
	}
}

func TestRockClient_ListQueryLambdaTagVersions(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	versions, err := rc.ListQueryLambdaTagVersions(ctx, "latest")
	require.NoError(t, err)

	for _, l := range versions {
		log.Printf("version: %s", *l.Version)
	}
}

func TestRockClient_ListQueryLambdaVersions(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	versions, err := rc.ListQueryLambdaVersions(ctx, "commons", "events")
	require.NoError(t, err)

	for _, l := range versions {
		log.Printf("version: %s", *l.Version)
	}
}

func TestRockClient_ListQueryLambdaTags(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	tags, err := rc.ListQueryLambdaTags(ctx)
	require.NoError(t, err)

	for _, tag := range tags {
		log.Printf("tag: %s", *tag.TagName)
	}
}

func TestRockClient_ListQueryLambdaTags_forQL(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	tags, err := rc.ListQueryLambdaTags(ctx, option.WithQueryLambda("commons", "events"))
	require.NoError(t, err)

	for _, tag := range tags {
		log.Printf("tag: %s", *tag.TagName)
	}
}
