// Test helpers for kafka and kafka-connect

package rockset_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

type kafkaConfig struct {
	topic           string
	integrationName string
	workspace       string
	collection      string
}

func testKafka(ctx context.Context, t *testing.T, rc *rockset.RockClient, kc kafkaConfig) {
	i, err := rc.CreateKafkaIntegration(ctx, kc.integrationName, option.WithKafkaDataFormat(option.KafkaFormatJSON),
		option.WithKafkaIntegrationTopic(kc.topic))
	require.NoError(t, err)

	u := fmt.Sprintf("https://%s", os.Getenv("ROCKSET_APISERVER"))

	cc := ConnectorConfig{
		Name:                        kc.integrationName,
		ConnectorClass:              "rockset.RocksetSinkConnector",
		TasksMax:                    2,
		Topics:                      kc.topic,
		RocksetTaskThreads:          2,
		RocksetApiserverURL:         u,
		RocksetIntegrationKey:       *i.Kafka.ConnectionString,
		Format:                      string(option.KafkaFormatJSON),
		KeyConverter:                "org.apache.kafka.connect.storage.StringConverter",
		ValueConverter:              "org.apache.kafka.connect.storage.StringConverter",
		KeyConverterSchemasEnable:   false,
		ValueConverterSchemasEnable: false,
	}
	err = configureKafkaConnect("http://localhost:8083/connectors", kc.integrationName, cc)
	require.NoError(t, err)

	t.Log("waiting for integration to be ready...")
	err = rc.WaitUntilKafkaIntegrationActive(ctx, kc.integrationName)
	require.NoError(t, err)

	_, err = rc.CreateKafkaCollection(ctx, kc.workspace, kc.collection,
		option.WithKafkaSource(kc.integrationName, kc.topic, option.KafkaStartingOffsetEarliest, option.WithJSONFormat()))
	require.NoError(t, err)

	t.Log("waiting for collection to start receiving documents...")
	err = rc.WaitUntilCollectionDocuments(ctx, kc.workspace, kc.collection, 1)
	require.NoError(t, err)

	t.Log("done")
}

type ConnectorRequest struct {
	Name   string          `json:"name"`
	Config ConnectorConfig `json:"config"`
}

type ConnectorConfig struct {
	Name                        string `json:"name"`
	ConnectorClass              string `json:"connector.class"`
	TasksMax                    int    `json:"tasks.max"`
	Topics                      string `json:"topics"`
	RocksetTaskThreads          int    `json:"rockset.task.threads"`
	RocksetApiserverURL         string `json:"rockset.apiserver.url"`
	RocksetIntegrationKey       string `json:"rockset.integration.key"`
	Format                      string `json:"format"`
	KeyConverter                string `json:"key.converter"`
	ValueConverter              string `json:"value.converter"`
	KeyConverterSchemasEnable   bool   `json:"key.converter.schemas.enable"`
	ValueConverterSchemasEnable bool   `json:"value.converter.schemas.enable"`
}

func loggingCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("failed to close: %v", err)
	}
}

func configureKafkaConnect(url, name string, cfg ConnectorConfig) error {
	r := ConnectorRequest{
		Name:   name,
		Config: cfg,
	}
	body, err := json.Marshal(r)
	if err != nil {
		return err
	}

	c := http.Client{}
	resp, err := c.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("payload: %s", string(payload))

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("bad response: %d", resp.StatusCode)
	}

	return err
}

func waitForKafkaConnect(t *testing.T, url string) func() error {
	return func() error {
		c := http.Client{}
		resp, err := c.Get(url)
		if err != nil {
			t.Logf("error getting %s: %v", url, err)
			return err
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Logf("error reading body: %v", err)
			return err
		}
		defer loggingCloser(resp.Body)

		if resp.StatusCode != 200 {
			t.Logf("error: %v", err)
			return fmt.Errorf("expected 200 got %d", resp.StatusCode)
		}

		t.Logf("body: %s", string(body))

		//type info struct {
		//	Version        string `json:"version"`
		//	Commit         string `json:"commit"`
		//	KafkaClusterID string `json:"kafka_cluster_id"`
		//}
		//var i info

		var i []string

		if err = json.Unmarshal(body, &i); err != nil {
			return err
		}

		t.Logf("%+v", i)

		return nil
	}
}

func environment(bootstrapServers, username, password string, format option.KafkaFormat) []string {
	env := []string{
		"CONNECT_GROUP_ID=rockset",
		"CONNECT_REST_ADVERTISED_HOST_NAME=rockset", // should be configurable
		"CONNECT_CONFIG_STORAGE_TOPIC=connect_config",
		"CONNECT_CONFIG_STORAGE_REPLICATION_FACTOR=3",
		"CONNECT_OFFSET_STORAGE_TOPIC=connect_offset",
		"CONNECT_OFFSET_FLUSH_INTERVAL_MS=10000",
		"CONNECT_OFFSET_STORAGE_FILE_FILENAME=/tmp/connect.offsets",
		"CONNECT_OFFSET_STORAGE_REPLICATION_FACTOR=3",
		"CONNECT_OFFSET_STORAGE_PARTITIONS=1",
		"CONNECT_STATUS_STORAGE_TOPIC=connect_status",
		"CONNECT_STATUS_STORAGE_REPLICATION_FACTOR=3",
	}

	switch format {
	case option.KafkaFormatJSON:
		env = append(env,
			"CONNECT_KEY_CONVERTER=org.apache.kafka.connect.json.JsonConverter",
			"CONNECT_KEY_CONVERTER_SCHEMAS_ENABLE=false",
			"CONNECT_VALUE_CONVERTER=org.apache.kafka.connect.json.JsonConverter",
			"CONNECT_VALUE_CONVERTER_SCHEMAS_ENABLE=false",
		)
	case option.KafkaFormatAVRO:
		fallthrough
	default:
		panic("not implemented")
	}

	env = append(env, connectParams("", bootstrapServers, username, password)...)
	env = append(env, connectParams("PRODUCER_", bootstrapServers, username, password)...)
	env = append(env, connectParams("CONSUMER_", bootstrapServers, username, password)...)

	return env
}

func connectParams(prefix, bootstrapServers, username, password string) []string {
	return []string{
		fmt.Sprintf("CONNECT_%sBOOTSTRAP_SERVERS=%s", prefix, bootstrapServers),
		fmt.Sprintf("CONNECT_%sSSL_ENDPOINT_IDENTIFICATION_ALGORITHM=https", prefix),
		fmt.Sprintf("CONNECT_%sSECURITY_PROTOCOL=SASL_SSL", prefix),
		fmt.Sprintf("CONNECT_%sSASL_MECHANISM=PLAIN", prefix),
		fmt.Sprintf(`CONNECT_%sSASL_JAAS_CONFIG=org.apache.kafka.common.security.plain.PlainLoginModule required username="%s" password="%s";`, prefix, username, password),
		fmt.Sprintf("CONNECT_%sREQUEST_TIMEOUT_MS=20000", prefix),
		fmt.Sprintf("CONNECT_%sRETRY_BACKOFF_MS=500", prefix),
	}
}
