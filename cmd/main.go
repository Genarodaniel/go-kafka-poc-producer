package main

import (
	"go-kafka-order-producer/config"
	"go-kafka-order-producer/internal/infra/events/kafka"
	"go-kafka-order-producer/internal/server"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	kafkaClient, err := kafka.NewKafka(config.Config.KafkaSeeds, config.Config.KafkaTopics)
	if err != nil {
		panic(err)
	}

	s := server.Init(&kafkaClient)
	s.Run()

}
