package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

func TestScheduledLambdas_CRUD(t *testing.T) {
	t.Parallel()

	ctx := test.Context()
	rc, randomName := vcrTestClient(t, t.Name())

    qlName := randomName("ql")
	keyName := randomName("key")
	ws := "commons"
	cronString := "0 0 0 ? * * *"
	var deleted bool

	key, err := rc.CreateAPIKey(ctx, keyName, option.WithRole("integration-test"))
	require.NoError(t, err)
	ql, err := rc.CreateQueryLambda(ctx, ws, qlName, "SELECT 1",
		option.WithQueryLambdaDescription("create"))
		require.NoError(t, err)
	qlVersion := ql.GetVersion()

	scheduledQl, err := rc.CreateScheduledLambda(ctx, ws, key.Key, cronString, qlName, option.WithScheduledLambdaVersion(qlVersion), option.WithScheduledLambdaTotalTimesToExecute(1))
	require.NoError(t, err)
	assert.Equal(t, int64(1), scheduledQl.GetTotalTimesToExecute())
	assert.Equal(t, qlName, scheduledQl.GetQlName())
	assert.Equal(t, cronString, scheduledQl.GetCronString())
	assert.Equal(t, qlVersion, scheduledQl.GetVersion())
	scheduledLambdaRRN := *scheduledQl.Rrn

	t.Cleanup(func() {
		if !deleted {
			err := rc.DeleteScheduledLambda(ctx, ws, scheduledLambdaRRN)
			require.NoError(t, err)
			err = rc.DeleteAPIKey(ctx, keyName)
			require.NoError(t, err)
			err = rc.DeleteQueryLambda(ctx, ws, qlName)
			require.NoError(t, err)
		}
	})

	err = rc.Wait.UntilScheduledLambdaAvailable(ctx, ws, scheduledLambdaRRN)
	require.NoError(t, err)

	getScheduledQl, err := rc.GetScheduledLambda(ctx, ws, scheduledLambdaRRN)
	require.NoError(t, err)
	assert.Equal(t, qlName, getScheduledQl.GetQlName())
	assert.Equal(t, cronString, getScheduledQl.GetCronString())
	assert.Equal(t, qlVersion, getScheduledQl.GetVersion())

	updateScheduledQl, err := rc.UpdateScheduledLambda(ctx, ws, scheduledLambdaRRN, option.WithScheduledLambdaTotalTimesToExecute(2))
	require.NoError(t, err)
	assert.Equal(t, int64(2), updateScheduledQl.GetTotalTimesToExecute())

	err = rc.Wait.UntilScheduledLambdaAvailable(ctx, ws, scheduledLambdaRRN)
	require.NoError(t, err)

	err = rc.DeleteScheduledLambda(ctx, ws, scheduledLambdaRRN)
	require.NoError(t, err)
	err = rc.DeleteAPIKey(ctx, keyName)
	require.NoError(t, err)
	err = rc.DeleteQueryLambda(ctx, ws, qlName)
	require.NoError(t, err)
	deleted = true
}
