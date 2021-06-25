package rockset_test

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func TestRockClient_ListUsers(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	users, err := rc.ListUsers(ctx)
	require.NoError(t, err)

	assert.NotEmpty(t, users)

	for _, user := range users {
		log.Debug().Str("user", user.Email).Send()
	}
}

func TestRockClient_GetCurrentUser(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	user, err := rc.GetCurrentUser(ctx)
	require.NoError(t, err)

	assert.Equal(t, "pme+circleci@rockset.com", user.Email)
}
