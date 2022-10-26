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
	suite.Run(s.T(), &s)
}

func (s *AliasIntegrationSuite) TestGetAlias() {
	skipUnlessIntegrationTest(s.T())

	ctx := testCtx()

	alias, err := s.rc.GetAlias(ctx, persistentWorkspace, persistentAlias)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "pme@rockset.com", alias.GetCreatorEmail())
}

func (s *AliasIntegrationSuite) TestListAliases() {
	skipUnlessIntegrationTest(s.T())

	ctx := testCtx()

	aliases, err := s.rc.ListAliases(ctx)
	require.NoError(s.T(), err)
	for _, a := range aliases {
		s.T().Logf("workspace: %s", a.GetName())
	}
}

func (s *AliasIntegrationSuite) TestListAliasesForWorkspace() {
	skipUnlessIntegrationTest(s.T())

	ctx := testCtx()

	aliases, err := s.rc.ListAliases(ctx, option.WithAliasWorkspace(persistentWorkspace))
	require.NoError(s.T(), err)
	for _, a := range aliases {
		s.T().Logf("workspace: %s", a.GetName())
	}
}

func (s *AliasIntegrationSuite) TestAliases() {
	skipUnlessIntegrationTest(s.T())

	ctx := testCtx()

	alias := randomName("alias")

	_, err := s.rc.CreateAlias(ctx, persistentWorkspace, alias, []string{"commons._events"})
	require.NoError(s.T(), err)

	err = s.rc.WaitUntilAliasAvailable(ctx, persistentWorkspace, alias)
	require.NoError(s.T(), err)

	// update

	err = s.rc.WaitUntilAliasAvailable(ctx, persistentWorkspace, alias)
	require.NoError(s.T(), err)

	err = s.rc.DeleteAlias(ctx, persistentWorkspace, alias)
	require.NoError(s.T(), err)
}
