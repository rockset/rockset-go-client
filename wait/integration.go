package wait

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/option"
)

// UntilKafkaIntegrationActive waits until all topics in a Kafka integration are in ACTIVE state.
func (w *Waiter) UntilKafkaIntegrationActive(ctx context.Context, integration string) error {
	return w.rc.RetryWithCheck(ctx, func() (bool, error) {
		zl := zerolog.Ctx(ctx)

		i, err := w.rc.GetIntegration(ctx, integration)
		if err != nil {
			return false, err
		}

		if i.Kafka == nil {
			return false, fmt.Errorf("not a kafka integration: %s", integration)
		}

		if i.Kafka.SourceStatusByTopic == nil {
			// for v3 integrations this will be nil
			zl.Trace().Bool("v3", i.Kafka.GetUseV3()).Msg("no topics found")
			return false, nil
		}

		var allActive = true
		for topic, status := range *i.Kafka.SourceStatusByTopic {
			zl.Trace().Str("state", status.GetState()).Str("topic", topic).Send()
			if status.GetState() != string(option.KafkaIntegrationActive) {
				allActive = false
			}
		}

		return !allActive, nil
	})
}
