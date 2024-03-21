package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

type SuiteScheduledLambda struct {
	suite.Suite
	rc      *rockset.RockClient
	qlName  string
	keyName string
	key     openapi.ApiKey
	qlVersion string
	ws        string
	scheduledLambdaRrn string
	cronString string
}

func TestSuiteScheduledLambdas(t *testing.T) {
	rc, nameGenerator := vcrTestClient(t, t.Name())

	s := SuiteScheduledLambda{
		rc:      rc,
		qlName: nameGenerator("ql"),
		keyName: nameGenerator("key"),
		ws:   "commons",
		cronString: "0 0 0 ? * * *",
	}
	suite.Run(t, &s)
}

func (s *SuiteScheduledLambda) SetupSuite() {
	ctx := test.Context()
	key, err := s.rc.CreateAPIKey(ctx, s.keyName, option.WithRole("integration-test"))
	s.Require().NoError(err)
	s.key = key
	ql, err := s.rc.CreateQueryLambda(ctx, s.ws, s.qlName, "SELECT 1",
		option.WithQueryLambdaDescription("create"))
	s.Require().NoError(err)
	s.qlVersion = ql.GetVersion()

	scheduledQl, err := s.rc.CreateScheduledLambda(ctx, s.ws, s.key.Key, s.cronString, s.qlName, option.WithScheduledLambdaVersion(s.qlVersion), option.WithScheduledLambdaTotalTimesToExecute(1))
	s.Require().NoError(err)
	s.Equal(int64(1), scheduledQl.GetTotalTimesToExecute())
	s.Equal(s.qlName, scheduledQl.GetQlName())
	s.Equal(s.cronString, scheduledQl.GetCronString())
	s.Equal(s.qlVersion, scheduledQl.GetVersion())
	s.scheduledLambdaRrn = *scheduledQl.Rrn
}

func (s *SuiteScheduledLambda) TearDownSuite() {
	ctx := test.Context()
	err := s.rc.DeleteScheduledLambda(ctx, s.ws, s.scheduledLambdaRrn)
	s.Require().NoError(err)
	err = s.rc.DeleteAPIKey(ctx, s.keyName)
	s.Require().NoError(err)
	err = s.rc.DeleteQueryLambda(ctx, s.ws, s.qlName)
	s.Require().NoError(err)
}

func (s *SuiteScheduledLambda) TestGetScheduledLambda() {
	ctx := test.Context()
    scheduledQl, err := s.rc.GetScheduledLambda(ctx, s.ws, s.scheduledLambdaRrn)
	s.Require().NoError(err)
	s.Equal(s.qlName, scheduledQl.GetQlName())
	s.Equal(s.cronString, scheduledQl.GetCronString())
	s.Equal(s.qlVersion, scheduledQl.GetVersion())
}

func (s *SuiteScheduledLambda) TestUpdateScheduledLambda() {
	ctx := test.Context()
	scheduledQl, err := s.rc.UpdateScheduledLambda(ctx, s.ws, s.scheduledLambdaRrn, option.WithScheduledLambdaTotalTimesToExecute(2))
	s.Require().NoError(err)
	s.Equal(int64(2), scheduledQl.GetTotalTimesToExecute())
}
