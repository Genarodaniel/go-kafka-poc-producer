package api

import (
	"go-kafka-order-producer/internal/api/healthcheck"
	"go-kafka-order-producer/internal/api/order"
	"go-kafka-order-producer/internal/infra/events/kafka"

	"github.com/gin-gonic/gin"
	"github.com/twmb/franz-go/pkg/kgo"
)

func Router(e *gin.Engine, kafkaClient *kgo.Client) {
	v1 := e.Group("/v1")

	orderGroup := v1.Group("/order")
	healthCheckGroup := v1.Group("/healthcheck")

	kafkaProducer := kafka.NewKafkaProducer(kafkaClient)
	healthcheck.Router(healthCheckGroup)
	order.Router(orderGroup, kafkaProducer)
}
