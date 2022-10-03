// Test suite for Confluent Cloud with a local kafka-connect
package rockset_test

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"
)

type CCKCTestSuite struct {
	suite.Suite
	rc         *rockset.RockClient
	dockerPool *dockertest.Pool
	connect    *dockertest.Resource
	kc         kafkaConfig
}

// Test creating an integration and collection for Confluent Cloud with a local kafka-connect
func TestCCKCSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := CCKCTestSuite{
		rc: rc,
		kc: kafkaConfig{
			topic:           "test_json",
			integrationName: "go-test",
			workspace:       "commons",
			collection:      "kafka-test",
		},
	}
	suite.Run(t, &s)
}

func (s *CCKCTestSuite) TestKafka() {
	ctx := testCtx()
	testKafka(ctx, s.T(), s.rc, s.kc)
}

func (s *CCKCTestSuite) SetupSuite() {
	var err error

	s.dockerPool, err = dockertest.NewPool("")
	s.Require().NoError(err)

	err = s.dockerPool.Client.Ping()
	s.Require().NoError(err)

	s.connect, err = s.dockerPool.RunWithOptions(&dockertest.RunOptions{
		Name:       "kafka-connect",
		Repository: "rockset/kafka-connect",
		Tag:        "1.4.2-5",
		Hostname:   "connect",
		Env: environment(os.Getenv("CC_BOOTSTRAP_SERVER"),
			os.Getenv("CC_KEY"),
			os.Getenv("CC_SECRET"),
			option.KafkaFormatJSON),
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9094/tcp": {{HostIP: "localhost", HostPort: "9094/tcp"}},
			"8083/tcp": {{HostIP: "localhost", HostPort: "8083/tcp"}},
		},
		ExposedPorts: []string{"9094/tcp", "8083/tcp"},
	})
	s.Require().NoError(err, "could not start kafka connect")

	s.dockerPool.MaxWait = 5 * time.Minute
	err = s.dockerPool.Retry(waitForKafkaConnect(s.T(), "http://localhost:8083"))
	s.Require().NoError(err, "could not start kafka connect")

}

func (s *CCKCTestSuite) TearDownSuite() {
	ctx := testCtx()

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

	// TODO: delete topics

	err = s.dockerPool.Purge(s.connect)
	s.Assert().NoError(err, "could not purge kafka connect")
}
