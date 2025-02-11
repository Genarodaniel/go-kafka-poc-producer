package main

import (
	"fmt"
	"go-kafka-order-producer/config"
	"go-kafka-order-producer/internal/server"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	seeds := []string{fmt.Sprintf("%s:%s", config.Config.KafkaHost, config.Config.KafkaPort)}
	fmt.Println(seeds)
	kafkaClient, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
	)
	if err != nil {
		panic(err)
	}
	defer kafkaClient.Close()

	s := server.Init(kafkaClient)
	s.Run()

}
