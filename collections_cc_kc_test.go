// Test suite for Confluent Cloud with a local kafka-connect

package rockset_test

import (
	"context"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ConfluentCloudWithKafkaConnectIntegrationSuite struct {
	suite.Suite
	rc               *rockset.RockClient
	dockerPool       *dockertest.Pool
	connect          *dockertest.Resource
	kc               kafkaConfig
	bootstrapServers string
	confluentKey     string
	confluentSecret  string
}

// Test creating an integration and collection for Confluent Cloud with a local kafka-connect
func TestConfluentCloudWithKafkaConnectIntegrationSuite(t *testing.T) {
	t.Skip("skipping kafka tests - too flakey :(")
	skipUnlessIntegrationTest(t)
	skipUnlessDocker(t)
	t.Parallel()

	name := randomName("cckc")

	s := ConfluentCloudWithKafkaConnectIntegrationSuite{
		rc: testClient(t),
		kc: kafkaConfig{
			topic:           "test_json",
			integrationName: name,
			workspace:       name,
			collection:      name,
		},
		bootstrapServers: skipUnlessEnvSet(t, "CC_BOOTSTRAP_SERVERS"),
		confluentKey:     skipUnlessEnvSet(t, "CC_KEY"),
		confluentSecret:  skipUnlessEnvSet(t, "CC_SECRET"),
	}
	suite.Run(t, &s)
}

func (s *ConfluentCloudWithKafkaConnectIntegrationSuite) TestKafka() {
	ctx := testCtx()
	c, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	testKafka(c, s.T(), s.rc, s.kc)
}

func (s *ConfluentCloudWithKafkaConnectIntegrationSuite) SetupSuite() {
	ctx := testCtx()
	var err error

	_, err = s.rc.CreateWorkspace(ctx, s.kc.workspace)
	s.Require().NoError(err)

	s.dockerPool, err = dockertest.NewPool("")
	s.Require().NoError(err)

	err = s.dockerPool.Client.Ping()
	s.Require().NoError(err)

	s.connect, err = s.dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "rockset/kafka-connect",
		Tag:        "latest",
		Hostname:   "connect",
		Env:        environment(s.bootstrapServers, s.confluentKey, s.confluentSecret, option.KafkaFormatJSON),
		// TODO don't use fixed host ports
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9094/tcp": {{HostIP: "localhost", HostPort: "9094/tcp"}},
			"8083/tcp": {{HostIP: "localhost", HostPort: "8083/tcp"}},
		},
		ExposedPorts: []string{"9094/tcp", "8083/tcp"},
	})
	s.Require().NoError(err, "could not start kafka connect")

	s.dockerPool.MaxWait = 5 * time.Minute
	err = s.dockerPool.Retry(waitForKafkaConnect(s.T(), "http://localhost:8083/connectors"))
	s.Require().NoError(err, "could not start kafka connect")

	s.T().Log("kafka connect running")
}

func (s *ConfluentCloudWithKafkaConnectIntegrationSuite) TearDownSuite() {
	ctx := testCtx()

	// TODO don't hardcode the URL
	if err := deleteConnector("http://localhost:8083/connectors", s.kc.integrationName); err == nil {
		s.T().Logf("deleted connector %s", s.kc.integrationName)
	}

	if err := s.rc.DeleteCollection(ctx, s.kc.workspace, s.kc.collection); err == nil {
		s.T().Logf("deleted collection %s.%s", s.kc.workspace, s.kc.collection)
	}

	err := s.rc.WaitUntilCollectionGone(ctx, s.kc.workspace, s.kc.collection)
	s.Require().NoError(err)

	if err := s.rc.DeleteIntegration(ctx, s.kc.integrationName); err == nil {
		s.T().Logf("deleted integration %s", s.kc.integrationName)
	} else {
		s.T().Logf("failed to delete integration %s: %v", s.kc.integrationName, err)
	}

	if err = s.rc.DeleteWorkspace(ctx, s.kc.workspace); err != nil {
		s.T().Logf("failed to delete workspace %s: %v", s.kc.workspace, err)
	}

	// TODO: delete topics

	err = s.dockerPool.Purge(s.connect)
	s.Assert().NoError(err, "could not purge kafka connect")
}
