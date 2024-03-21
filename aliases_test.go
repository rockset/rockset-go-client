package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

func TestAliases_getAlias(t *testing.T) {
	t.Parallel()

	ctx := test.Context()
	rc, _ := vcrTestClient(t, t.Name())

	alias, err := rc.GetAlias(ctx, test.PersistentWorkspace, test.PersistentAlias)
	require.NoError(t, err)
	assert.Equal(t, "pme@rockset.com", alias.GetCreatorEmail())
}

func TestAliases_listAliases(t *testing.T) {
	t.Parallel()

	ctx := test.Context()
	rc, _ := vcrTestClient(t, t.Name())

	aliases, err := rc.ListAliases(ctx)
	require.NoError(t, err)
	for _, a := range aliases {
		t.Logf("workspace: %s", a.GetName())
	}
}

func TestAliases_listAliasesForWorkspace(t *testing.T) {
	t.Parallel()

	ctx := test.Context()
	rc, _ := vcrTestClient(t, t.Name())

	aliases, err := rc.ListAliases(ctx, option.WithAliasWorkspace(test.PersistentWorkspace))
	require.NoError(t, err)
	for _, a := range aliases {
		t.Logf("workspace: %s", a.GetName())
	}
}

func TestAliases_CRUD(t *testing.T) {
	t.Parallel()

	ctx := test.Context()
	rc, randomName := vcrTestClient(t, t.Name())
	alias := randomName("alias")
	var deleted bool

	_, err := rc.CreateAlias(ctx, test.PersistentWorkspace, alias, []string{"commons._events"},
		option.WithAliasDescription("original"))
	require.NoError(t, err)

	t.Cleanup(func() {
		if !deleted {
			_ = rc.DeleteAlias(ctx, test.PersistentWorkspace, alias)
		}
	})

	err = rc.Wait.UntilAliasAvailable(ctx, test.PersistentWorkspace, alias)
	require.NoError(t, err)

	err = rc.UpdateAlias(ctx, test.PersistentWorkspace, alias, []string{"commons._events"},
		option.WithAliasDescription("updated"))
	require.NoError(t, err)

	err = rc.Wait.UntilAliasAvailable(ctx, test.PersistentWorkspace, alias)
	require.NoError(t, err)

	err = rc.DeleteAlias(ctx, test.PersistentWorkspace, alias)
	require.NoError(t, err)
	deleted = true
}
