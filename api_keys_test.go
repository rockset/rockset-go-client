package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

type SuiteAPIKey struct {
	suite.Suite
	rc      *rockset.RockClient
	keyName string
	key     openapi.ApiKey
}

func TestSuiteAPIKey(t *testing.T) {
	rc, randomName := vcrClient(t)

	s := SuiteAPIKey{
		rc:      rc,
		keyName: randomName("key"),
	}
	suite.Run(t, &s)
}

// this role is persistent and should always exist
const IntegrationTestRole = "integration-test"

func (s *SuiteAPIKey) SetupSuite() {
	ctx := testCtx()
	key, err := s.rc.CreateAPIKey(ctx, s.keyName, option.WithRole(IntegrationTestRole))
	s.Require().NoError(err)
	s.key = key
}

func (s *SuiteAPIKey) TearDownSuite() {
	ctx := testCtx()
	err := s.rc.DeleteAPIKey(ctx, s.keyName)
	s.Require().NoError(err)
}

const mask = "************************************************************"

func (s *SuiteAPIKey) TestGetAPIKey() {
	ctx := testCtx()
	key, err := s.rc.GetAPIKey(ctx, s.keyName)
	s.Require().NoError(err)
	s.Assert().Equalf(mask, key.Key[4:], "key should be masked")
}

func (s *SuiteAPIKey) TestUpdateAPIKey() {
	ctx := testCtx()
	key, err := s.rc.UpdateAPIKey(ctx, s.keyName, option.State(option.KeySuspended))
	s.Require().NoError(err)
	s.Assert().NotEqual(option.KeyActive, key.GetState())
}

func (s *SuiteAPIKey) TestGetAPIKeyWithReveal() {
	ctx := testCtx()
	key, err := s.rc.GetAPIKey(ctx, s.keyName, option.RevealKey())
	s.Require().NoError(err)
	s.Assert().NotEqualf(mask, key.Key[4:], "key should not be masked")
}

func (s *SuiteAPIKey) TestListAPIKeys() {
	ctx := testCtx()
	keys, err := s.rc.ListAPIKeys(ctx)
	s.Require().NoError(err)

	var found bool
	for _, key := range keys {
		if key.Name == s.keyName {
			found = true
		}
	}

	s.Assert().True(found)
}
