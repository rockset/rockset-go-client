// Test suite for a pure Confluent Cloud setup

package rockset_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/wait"
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
	test.SkipUnlessIntegrationTest(t)

	s := ConfluentCloudIntegrationSuite{
		rc:               test.Client(t),
		integrationName:  test.RandomName("integration"),
		ws:               test.RandomName("cc"),
		coll:             test.RandomName("cc"),
		topic:            "test_json",
		bootstrapServers: test.SkipUnlessEnvSet(t, "CC_BOOTSTRAP_SERVERS"),
		confluentKey:     test.SkipUnlessEnvSet(t, "CC_KEY"),
		confluentSecret:  test.SkipUnlessEnvSet(t, "CC_SECRET"),
	}
	suite.Run(t, &s)
}

func (s *ConfluentCloudIntegrationSuite) SetupSuite() {
	ctx := test.Context()

	_, err := s.rc.CreateWorkspace(ctx, s.ws)
	s.Require().NoError(err)

	_, err = s.rc.CreateKafkaIntegration(ctx, s.integrationName,
		option.WithKafkaIntegrationDescription(test.Description()),
		option.WithKafkaV3(),
		option.WithKafkaBootstrapServers(s.bootstrapServers),
		option.WithKafkaSecurityConfig(s.confluentKey, s.confluentSecret),
	)
	s.Require().NoError(err)

	err = s.rc.WaitUntilKafkaIntegrationActive(ctx, s.integrationName)
	s.Require().NoError(err)
}

func (s *ConfluentCloudIntegrationSuite) TearDownSuite() {
	ctx := test.Context()

	err := s.rc.DeleteIntegration(ctx, s.integrationName)
	s.NoError(err)

	err = s.rc.DeleteWorkspace(ctx, s.ws)
	s.NoError(err)
}

func (s *ConfluentCloudIntegrationSuite) TestCreateJSONCollection() {
	ctx := test.Context()

	_, err := s.rc.CreateKafkaCollection(ctx, s.ws, s.coll,
		option.WithCollectionDescription(test.Description()),
		option.WithCollectionRetention(time.Hour),
		option.WithKafkaSource(s.integrationName, s.topic, option.KafkaStartingOffsetEarliest, option.WithJSONFormat(),
			option.WithKafkaSourceV3(),
		),
	)
	s.Require().NoError(err)

	w := wait.New(s.rc)
	err = w.UntilCollectionReady(ctx, s.ws, s.coll)
	s.Require().NoError(err)

	// TODO(pmenglund) this should write a document to kafka so we don't need a data generator
	//  in Confluent Cloud
	err = w.UntilCollectionHasNewDocuments(ctx, s.ws, s.coll, 1)
	s.Require().NoError(err)
}

func (s *ConfluentCloudIntegrationSuite) TestDeleteCollection() {
	ctx := test.Context()

	err := s.rc.DeleteCollection(ctx, s.ws, s.coll)
	s.Require().NoError(err)

	w := wait.New(s.rc)
	err = w.UntilCollectionGone(ctx, s.ws, s.coll)
	s.Require().NoError(err)
}
