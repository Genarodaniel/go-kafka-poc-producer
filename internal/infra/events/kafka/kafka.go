package kafka

import (
	"context"
	"encoding/json"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaInterface interface {
	Produce(ctx context.Context, topic string, body any) error
}

type Kafka struct {
	Client *kgo.Client
}

func NewKafkaProducer(client *kgo.Client) KafkaInterface {
	return &Kafka{
		Client: client,
	}
}

func (k *Kafka) Produce(ctx context.Context, topic string, body any) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	record := &kgo.Record{Topic: topic, Value: payload}
	k.Client.ProduceSync(ctx, record)
	return nil
}

func (k *Kafka) TransformPayload(body any) ([]byte, error) {
	return json.Marshal(body)
}
