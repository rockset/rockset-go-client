// Test suite for self-managed kafka

package rockset_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-zookeeper/zk"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

type KafkaIntegrationSuite struct {
	suite.Suite
	rc               *rockset.RockClient
	dockerPool       *dockertest.Pool
	zookeeper        *dockertest.Resource
	kafka            *dockertest.Resource
	connect          *dockertest.Resource
	network          *docker.Network
	kc               kafkaConfig
	bootstrapServers string
	confluentKey     string
	confluentSecret  string
}

// Test creating an integration and collection for a self-managed kafka with local kafka-connect
func TestKafkaIntegrationSuite(t *testing.T) {
	t.Skip("skipping kafka tests - too flakey :(")
	test.SkipUnlessIntegrationTest(t)
	test.SkipUnlessDocker(t)

	s := KafkaIntegrationSuite{
		rc: test.Client(t),
		kc: kafkaConfig{
			topic:           "test_json",
			integrationName: test.RandomName("kafka"),
			workspace:       test.RandomName("kafka"),
			collection:      test.RandomName("kafka"),
		},
		bootstrapServers: test.SkipUnlessEnvSet(t, "CC_BOOTSTRAP_SERVERS"),
		confluentKey:     test.SkipUnlessEnvSet(t, "CC_KEY"),
		confluentSecret:  test.SkipUnlessEnvSet(t, "CC_SECRET"),
	}
	suite.Run(t, &s)
}

func (s *KafkaIntegrationSuite) TestKafka() {
	ctx := test.Context()
	testKafka(ctx, s.T(), s.rc, s.kc)
}

//nolint:funlen
func (s *KafkaIntegrationSuite) SetupSuite() {
	var err error
	ctx := test.Context()

	_, err = s.rc.CreateWorkspace(ctx, s.kc.workspace)
	s.Require().NoError(err)
	s.dockerPool, err = dockertest.NewPool("")
	s.Require().NoError(err)

	err = s.dockerPool.Client.Ping()
	s.Require().NoError(err)

	s.network, err = s.dockerPool.Client.CreateNetwork(docker.CreateNetworkOptions{Name: "zookeeper_kafka_network"})
	s.Require().NoError(err, "could not create a network to zookeeper and kafka")

	s.zookeeper, err = s.dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "wurstmeister/zookeeper",
		Tag:          "3.4.6",
		NetworkID:    s.network.ID,
		Hostname:     "zookeeper",
		ExposedPorts: []string{"2181"},
	})
	s.Require().NoError(err, "could not start zookeeper")

	conn, _, err := zk.Connect([]string{fmt.Sprintf("127.0.0.1:%s", s.zookeeper.GetPort("2181/tcp"))}, 10*time.Second)
	s.Require().NoError(err, "could not connect zookeeper")
	defer conn.Close()

	//nolint:exhaustive
	retryFn := func() error {
		switch conn.State() {
		case zk.StateHasSession, zk.StateConnected:
			return nil
		default:
			return errors.New("not yet connected")
		}
	}

	err = s.dockerPool.Retry(retryFn)
	s.Require().NoError(err, "could not connect to zookeeper")

	s.kafka, err = s.dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "wurstmeister/kafka",
		Tag:        "2.13-2.8.1",
		NetworkID:  s.network.ID,
		Hostname:   "kafka",
		Env: []string{
			"KAFKA_CREATE_TOPICS=domain.test:1:1:compact",
			"KAFKA_ADVERTISED_LISTENERS=INSIDE://kafka:9092,OUTSIDE://localhost:9093",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=INSIDE:PLAINTEXT,OUTSIDE:PLAINTEXT",
			"KAFKA_LISTENERS=INSIDE://0.0.0.0:9092,OUTSIDE://0.0.0.0:9093",
			"KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181",
			"KAFKA_INTER_BROKER_LISTENER_NAME=INSIDE",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9093/tcp": {{HostIP: "localhost", HostPort: "9093/tcp"}},
		},
		ExposedPorts: []string{"9093/tcp"},
	})
	s.Require().NoError(err, "could not start kafka")

	bootstrapServers := "localhost:9093"
	retryFn = func() error {
		deliveryChan := make(chan kafka.Event)
		producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": bootstrapServers,
			"acks":              "all",
		})
		if err != nil {
			return err
		}
		defer producer.Close()

		topic := "domain.test"
		message := &kafka.Message{
			Key: []byte("any-key"),
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte("Hello World"),
		}
		if err = producer.Produce(message, deliveryChan); err != nil {
			return err
		}
		s.T().Log("message produced")

		e := <-deliveryChan
		if e.(*kafka.Message).TopicPartition.Error != nil {
			return e.(*kafka.Message).TopicPartition.Error
		}
		s.T().Log("message consumed")

		return nil
	}

	err = s.dockerPool.Retry(retryFn)
	s.Require().NoError(err, "could not connect to kafka")

	// kafka connect

	s.connect, err = s.dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "rockset/kafka-connect",
		Tag:        "latest",
		Hostname:   "connect",
		Env:        environment(s.bootstrapServers, s.confluentKey, s.confluentSecret, option.KafkaFormatJSON),
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

	s.T().Log("all containers running")
}

func (s *KafkaIntegrationSuite) TearDownSuite() {
	ctx := test.Context()
	var err error

	if err = s.rc.DeleteCollection(ctx, s.kc.workspace, s.kc.collection); err == nil {
		s.T().Logf("deleted collection %s.%s", s.kc.workspace, s.kc.collection)
	}

	err = s.rc.Wait.UntilCollectionGone(ctx, s.kc.workspace, s.kc.collection)
	s.Require().NoError(err)

	if err := s.rc.DeleteIntegration(ctx, s.kc.integrationName); err == nil {
		s.T().Logf("deleted integration %s", s.kc.integrationName)
	} else {
		s.T().Logf("failed to delete integration %s: %v", s.kc.integrationName, err)
	}

	err = s.dockerPool.Purge(s.connect)
	s.Assert().NoError(err, "could not purge kafka connect")

	err = s.dockerPool.Purge(s.kafka)
	s.Assert().NoError(err, "could not purge kafka")

	err = s.dockerPool.Purge(s.zookeeper)
	s.Assert().NoError(err, "could not purge zookeeper")

	// sleep a little before removing the network so all containers have time to exit
	time.Sleep(5 * time.Second)

	err = s.dockerPool.Client.RemoveNetwork(s.network.ID)
	s.Assert().NoError(err, "could not remove network")

	err = s.rc.DeleteWorkspace(ctx, s.kc.workspace)
	s.NoError(err, "could not remove workspace %s", s.kc.workspace)
}
