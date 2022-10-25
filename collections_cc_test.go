// Test suite for a pure Confluent Cloud setup

package rockset_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type ConfluentCloudIntegrationSuite struct {
	suite.Suite
	rc               *rockset.RockClient
	integrationName  string
	ws               string
	coll             string
	topic            string
	bootstrapServers string
	confluentKey     string
	confluentSecret  string
}

// Test creating an integration and collection for Confluent Cloud
func TestConfluentCloudIntegrationSuite(t *testing.T) {
	t.Skip("skipping kafka tests - too flakey :(")
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := ConfluentCloudIntegrationSuite{
		rc:               rc,
		integrationName:  randomName(t, "integration"),
		ws:               "acc",
		coll:             randomName(t, "cc"),
		topic:            "test_json",
		bootstrapServers: skipUnlessEnvSet(t, "CC_BOOTSTRAP_SERVERS"),
		confluentKey:     skipUnlessEnvSet(t, "CC_KEY"),
		confluentSecret:  skipUnlessEnvSet(t, "CC_SECRET"),
	}
	suite.Run(t, &s)
}

func (s *ConfluentCloudIntegrationSuite) SetupSuite() {
	ctx := testCtx()

	_, err := s.rc.CreateKafkaIntegration(ctx, s.integrationName,
		option.WithKafkaIntegrationDescription(description()),
		option.WithKafkaV3(),
		option.WithKafkaBootstrapServers(s.bootstrapServers),
		option.WithKafkaSecurityConfig(s.confluentKey, s.confluentSecret),
	)
	s.Require().NoError(err)

	err = s.rc.WaitUntilKafkaIntegrationActive(ctx, s.integrationName)
	s.Require().NoError(err)
}

func (s *ConfluentCloudIntegrationSuite) TearDownSuite() {
	ctx := testCtx()

	err := s.rc.DeleteIntegration(ctx, s.integrationName)
	s.Require().NoError(err)
}

func (s *ConfluentCloudIntegrationSuite) TestCreateJSONCollection() {
	ctx := testCtx()

	_, err := s.rc.CreateKafkaCollection(ctx, s.ws, s.coll,
		option.WithCollectionDescription(description()),
		option.WithCollectionRetention(time.Hour),
		option.WithKafkaSource(s.integrationName, s.topic, option.KafkaStartingOffsetEarliest, option.WithJSONFormat(),
			option.WithKafkaSourceV3(),
		),
	)
	s.Require().NoError(err)

	err = s.rc.WaitUntilCollectionReady(ctx, s.ws, s.coll)
	s.Require().NoError(err)

	// TODO(pmenglund) this should write a document to kafka so we don't need a data generator
	//  in Confluent Cloud
	err = s.rc.WaitUntilCollectionHasNewDocuments(ctx, s.ws, s.coll, 1)
	s.Require().NoError(err)
}

func (s *ConfluentCloudIntegrationSuite) TestDeleteCollection() {
	ctx := testCtx()

	err := s.rc.DeleteCollection(ctx, s.ws, s.coll)
	s.Require().NoError(err)

	err = s.rc.WaitUntilCollectionGone(ctx, s.ws, s.coll)
	s.Require().NoError(err)
}
