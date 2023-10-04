package rockset_test

import (
	"github.com/rockset/rockset-go-client/internal/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type AliasIntegrationSuite struct {
	suite.Suite
	rc    *rockset.RockClient
	alias string
}

func TestAliasIntegrationSuite(t *testing.T) {
	rc, randomName := vcrTestClient(t, t.Name())
	s := AliasIntegrationSuite{rc: rc, alias: randomName("alias")}
	suite.Run(t, &s)
}

func (s *AliasIntegrationSuite) TestGetAlias() {
	ctx := test.Context()

	alias, err := s.rc.GetAlias(ctx, test.PersistentWorkspace, test.PersistentAlias)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "pme@rockset.com", alias.GetCreatorEmail())
}

func (s *AliasIntegrationSuite) TestListAliases() {
	ctx := test.Context()

	aliases, err := s.rc.ListAliases(ctx)
	require.NoError(s.T(), err)
	for _, a := range aliases {
		s.T().Logf("workspace: %s", a.GetName())
	}
}

func (s *AliasIntegrationSuite) TestListAliasesForWorkspace() {
	ctx := test.Context()

	aliases, err := s.rc.ListAliases(ctx, option.WithAliasWorkspace(test.PersistentWorkspace))
	require.NoError(s.T(), err)
	for _, a := range aliases {
		s.T().Logf("workspace: %s", a.GetName())
	}
}

func (s *AliasIntegrationSuite) TestAliases() {
	ctx := test.Context()

	_, err := s.rc.CreateAlias(ctx, test.PersistentWorkspace, s.alias, []string{"commons._events"})
	require.NoError(s.T(), err)

	err = s.rc.Wait.UntilAliasAvailable(ctx, test.PersistentWorkspace, s.alias)
	require.NoError(s.T(), err)

	// update

	err = s.rc.Wait.UntilAliasAvailable(ctx, test.PersistentWorkspace, s.alias)
	require.NoError(s.T(), err)

	err = s.rc.DeleteAlias(ctx, test.PersistentWorkspace, s.alias)
	require.NoError(s.T(), err)
}
