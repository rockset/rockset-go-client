package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
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
	rc, randomName := vcrTestClient(t, t.Name())

	s := SuiteAPIKey{
		rc:      rc,
		keyName: randomName("key"),
	}
	suite.Run(t, &s)
}

// this role is persistent and should always exist
const IntegrationTestRole = "integration-test"

func (s *SuiteAPIKey) SetupSuite() {
	ctx := test.Context()
	key, err := s.rc.CreateAPIKey(ctx, s.keyName, option.WithRole(IntegrationTestRole))
	s.Require().NoError(err)
	s.key = key
}

func (s *SuiteAPIKey) TearDownSuite() {
	ctx := test.Context()
	err := s.rc.DeleteAPIKey(ctx, s.keyName)
	s.Require().NoError(err)
}

func (s *SuiteAPIKey) TestGetAPIKey() {
	ctx := test.Context()
	key, err := s.rc.GetAPIKey(ctx, s.keyName)
	s.Require().NoError(err)
	s.Equal(6, len(key.GetKey()), "key hash should be 6 characters")
	key.GetCreatedByApikeyName()
	key.GetLastAccessTime()
}

func (s *SuiteAPIKey) TestUpdateAPIKey() {
	ctx := test.Context()
	key, err := s.rc.UpdateAPIKey(ctx, s.keyName, option.State(option.KeySuspended))
	s.Require().NoError(err)
	s.NotEqual(option.KeyActive, key.GetState())
}

func (s *SuiteAPIKey) TestListAPIKeys() {
	ctx := test.Context()
	keys, err := s.rc.ListAPIKeys(ctx)
	s.Require().NoError(err)

	var found bool
	for _, key := range keys {
		if key.Name == s.keyName {
			found = true
		}
	}

	s.True(found)
}
