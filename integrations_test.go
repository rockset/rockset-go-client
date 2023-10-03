package rockset_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

type IntegrationTestSuite struct {
	suite.Suite
	rc             *rockset.RockClient
	s3Integration  string
	gcsIntegration string
}

func TestIntegrationTestSuite(t *testing.T) {
	_, randomName := vcrTestClient(t, t.Name())

	suite.Run(t, &IntegrationTestSuite{
		s3Integration:  randomName("s3"),
		gcsIntegration: randomName("gcs"),
	})
}

func (s *IntegrationTestSuite) BeforeTest(suiteName, testName string) {
	s.rc, _ = vcrTestClient(s.T(), fmt.Sprintf("%s/%s", suiteName, testName))
}

func (s *IntegrationTestSuite) TearDown() {
	ctx := test.Context()

	// clean up any lingering integrations
	_ = s.rc.DeleteIntegration(ctx, s.s3Integration)
	_ = s.rc.DeleteIntegration(ctx, s.gcsIntegration)
}

func (s *IntegrationTestSuite) TestGetIntegration() {
	ctx := test.Context()

	const iName = "acc-kafka-integration"
	integration, err := s.rc.GetIntegration(ctx, iName)
	s.NoError(err)
	s.Assert().Equal(iName, integration.GetName())
}

func (s *IntegrationTestSuite) TestListIntegrations() {
	ctx := test.Context()

	_, err := s.rc.ListIntegrations(ctx)
	s.NoError(err)
}

func (s *IntegrationTestSuite) TestCreateS3Integration() {
	ctx := test.Context()

	_, err := s.rc.CreateS3Integration(ctx, s.s3Integration,
		option.AWSRole("arn:aws:iam::469279130686:role/rockset-s3-integration"),
		option.WithS3IntegrationDescription(test.Description()))
	s.Require().NoError(err)

	err = s.rc.DeleteIntegration(ctx, s.s3Integration)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestCreateGCSIntegration() {
	saKeyFile := test.SkipUnlessEnvSet(s.T(), "GCP_SA_KEY_JSON")

	ctx := test.Context()

	_, err := s.rc.CreateGCSIntegration(ctx, s.gcsIntegration, saKeyFile,
		option.WithGCSIntegrationDescription(test.Description()))
	s.Require().NoError(err)

	err = s.rc.DeleteIntegration(ctx, s.gcsIntegration)
	s.Require().NoError(err)
}
