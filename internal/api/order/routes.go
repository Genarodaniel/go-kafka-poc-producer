package order

import (
	"go-kafka-order-producer/internal/infra/events/kafka"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, kafkaProducer kafka.KafkaInterface) {
	service := NewOrderService(kafkaProducer)
	handler := NewOrderHandler(service)

	g.POST("/", handler.HandlePostOrder)
}
