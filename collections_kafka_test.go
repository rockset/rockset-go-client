package rockset_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type SuiteKafkaCollection struct {
	suite.Suite
	rc    *rockset.RockClient
	name  string
	ws    string
	coll  string
	topic string
}

func TestSuiteKafkaCollection(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := SuiteKafkaCollection{
		rc:    rc,
		name:  "confluent-cloud-unit-test",
		ws:    "tests",
		coll:  "confluent-cloud-unit-test",
		topic: "test_json",
	}
	suite.Run(t, &s)
}

func (s *SuiteKafkaCollection) SetupSuite() {
	apikey := skipUnlessEnvSet(s.T(), "CC_KEY")
	secret := skipUnlessEnvSet(s.T(), "CC_SECRET")
	bootstrap := skipUnlessEnvSet(s.T(), "CC_BOOTSTRAP_SERVER")
	ctx := testCtx()

	_, err := s.rc.CreateKafkaIntegration(ctx, s.name,
		option.WithKafkaIntegrationDescription("created by go integration test"),
		option.WithKafkaV3(),
		option.WithKafkaBootstrapServers(bootstrap),
		option.WithKafkaSecurityConfig(apikey, secret),
	)
	s.Require().NoError(err)
}

func (s *SuiteKafkaCollection) TearDownSuite() {
	ctx := testCtx()

	err := s.rc.DeleteIntegration(ctx, s.name)
	s.Require().NoError(err)
}

func (s *SuiteKafkaCollection) TestCreateJSONCollection() {
	ctx := testCtx()

	_, err := s.rc.CreateKafkaCollection(ctx, s.ws, s.coll,
		option.WithCollectionDescription("created by go integration test"),
		option.WithCollectionRetention(time.Hour),
		option.WithKafkaSource(s.name, s.topic, option.KafkaStartingOffsetEarliest, option.WithJSONFormat(),
			option.WithKafkaSourceV3(),
		),
	)
	s.Require().NoError(err)

	err = s.rc.WaitUntilCollectionReady(ctx, s.ws, s.name)
	s.Require().NoError(err)

	// TODO(pmenglund) this should write a document to kafka so we don't need a data generator
	//  in Confluent Cloud
	err = s.rc.WaitUntilCollectionDocuments(ctx, s.ws, s.coll, 1)
	s.Require().NoError(err)
}

func (s *SuiteKafkaCollection) TestDeleteCollection() {
	ctx := testCtx()

	err := s.rc.DeleteCollection(ctx, s.ws, s.name)
	s.Require().NoError(err)

	err = s.rc.WaitUntilCollectionGone(ctx, s.ws, s.coll)
	s.Require().NoError(err)
}
