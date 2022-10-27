package rockset_test

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type IntegrationTestSuite struct {
	suite.Suite
	rc             *rockset.RockClient
	s3Integration  string
	gcsIntegration string
}

func TestIntegrationTestSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)

	suite.Run(t, &IntegrationTestSuite{
		rc:             testClient(t),
		s3Integration:  randomName("s3"),
		gcsIntegration: randomName("gcs"),
	})
}

func (s *IntegrationTestSuite) TearDown() {
	ctx := testCtx()

	// clean up any lingering integrations
	_ = s.rc.DeleteIntegration(ctx, s.s3Integration)
	_ = s.rc.DeleteIntegration(ctx, s.gcsIntegration)
}

func (s *IntegrationTestSuite) TestGetIntegration() {
	ctx := testCtx()

	const iName = "confluent-cloud"
	integration, err := s.rc.GetIntegration(ctx, iName)
	s.NoError(err)
	s.Assert().Equal(iName, integration.GetName())
}

func (s *IntegrationTestSuite) TestListIntegrations() {
	ctx := testCtx()

	_, err := s.rc.ListIntegrations(ctx)
	s.NoError(err)
}

func (s *IntegrationTestSuite) TestCreateS3Integration() {
	ctx := testCtx()

	_, err := s.rc.CreateS3Integration(ctx, s.s3Integration,
		option.AWSRole("arn:aws:iam::469279130686:role/rockset-s3-integration"),
		option.WithS3IntegrationDescription(description()))
	s.Require().NoError(err)

	err = s.rc.DeleteIntegration(ctx, s.s3Integration)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestCreateGCSIntegration() {
	saKeyFile := skipUnlessEnvSet(s.T(), "GCP_SA_KEY_JSON")

	ctx := testCtx()

	_, err := s.rc.CreateGCSIntegration(ctx, s.gcsIntegration, saKeyFile,
		option.WithGCSIntegrationDescription(description()))
	s.Require().NoError(err)

	err = s.rc.DeleteIntegration(ctx, s.gcsIntegration)
	s.Require().NoError(err)
}
