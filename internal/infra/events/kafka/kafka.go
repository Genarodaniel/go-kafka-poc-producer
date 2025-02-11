package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaInterface interface {
	Produce(ctx context.Context, topic string, key string, body any) error
}

type Kafka struct {
	Client *kgo.Client
}

func NewKafka(seeds []string, topics []string) (KafkaInterface, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumeTopics(topics...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka client: %s", err.Error())
	}

	return &Kafka{
		Client: client,
	}, nil
}

func (k *Kafka) Produce(ctx context.Context, topic string, key string, body any) error {
	payload, err := k.SerializePayload(body)
	if err != nil {
		return err
	}

	record := &kgo.Record{Topic: topic, Value: payload, Key: []byte(key)}
	k.Client.ProduceSync(ctx, record)
	return nil
}

func (k *Kafka) SerializePayload(body any) ([]byte, error) {
	response, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("Error while serializing payload %s", err.Error())
	}
	return response, nil
}
