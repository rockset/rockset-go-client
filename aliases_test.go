package rockset_test

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type AliasIntegrationSuite struct {
	suite.Suite
	rc *rockset.RockClient
}

func TestAliasIntegrationSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := AliasIntegrationSuite{rc: rc}
	suite.Run(t, &s)
}

func (s *AliasIntegrationSuite) TestGetAlias(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	alias, err := s.rc.GetAlias(ctx, persistentWorkspace, persistentAlias)
	require.NoError(t, err)
	assert.Equal(t, "pme@rockset.com", alias.GetCreatorEmail())
}

func (s *AliasIntegrationSuite) TestListAliases(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	aliases, err := s.rc.ListAliases(ctx)
	require.NoError(t, err)
	for _, a := range aliases {
		t.Logf("workspace: %s", a.GetName())
	}
}

func (s *AliasIntegrationSuite) TestListAliasesForWorkspace(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	aliases, err := s.rc.ListAliases(ctx, option.WithAliasWorkspace(persistentWorkspace))
	require.NoError(t, err)
	for _, a := range aliases {
		t.Logf("workspace: %s", a.GetName())
	}
}

func (s *AliasIntegrationSuite) TestAliases(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	alias := randomName("alias")

	_, err := s.rc.CreateAlias(ctx, persistentWorkspace, alias, []string{"commons._events"})
	require.NoError(t, err)

	err = s.rc.WaitUntilAliasAvailable(ctx, persistentWorkspace, alias)
	require.NoError(t, err)

	// update

	err = s.rc.WaitUntilAliasAvailable(ctx, persistentWorkspace, alias)
	require.NoError(t, err)

	err = s.rc.DeleteAlias(ctx, persistentWorkspace, alias)
	require.NoError(t, err)
}
