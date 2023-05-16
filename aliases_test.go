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
	rc    *rockset.RockClient
	alias string
}

func TestAliasIntegrationSuite(t *testing.T) {
	rc, randomName := vcrClient(t, t.Name())
	s := AliasIntegrationSuite{rc: rc, alias: randomName("alias")}
	suite.Run(t, &s)
}

func (s *AliasIntegrationSuite) TestGetAlias() {
	ctx := testCtx()

	alias, err := s.rc.GetAlias(ctx, persistentWorkspace, persistentAlias)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "pme@rockset.com", alias.GetCreatorEmail())
}

func (s *AliasIntegrationSuite) TestListAliases() {
	ctx := testCtx()

	aliases, err := s.rc.ListAliases(ctx)
	require.NoError(s.T(), err)
	for _, a := range aliases {
		s.T().Logf("workspace: %s", a.GetName())
	}
}

func (s *AliasIntegrationSuite) TestListAliasesForWorkspace() {
	ctx := testCtx()

	aliases, err := s.rc.ListAliases(ctx, option.WithAliasWorkspace(persistentWorkspace))
	require.NoError(s.T(), err)
	for _, a := range aliases {
		s.T().Logf("workspace: %s", a.GetName())
	}
}

func (s *AliasIntegrationSuite) TestAliases() {
	ctx := testCtx()

	_, err := s.rc.CreateAlias(ctx, persistentWorkspace, s.alias, []string{"commons._events"})
	require.NoError(s.T(), err)

	err = s.rc.WaitUntilAliasAvailable(ctx, persistentWorkspace, s.alias)
	require.NoError(s.T(), err)

	// update

	err = s.rc.WaitUntilAliasAvailable(ctx, persistentWorkspace, s.alias)
	require.NoError(s.T(), err)

	err = s.rc.DeleteAlias(ctx, persistentWorkspace, s.alias)
	require.NoError(s.T(), err)
}
