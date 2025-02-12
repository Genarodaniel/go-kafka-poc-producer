package server

import (
	"go-kafka-order-producer/internal/api"
	"go-kafka-order-producer/internal/infra/events/kafka"

	"github.com/gin-gonic/gin"
)

func Init(kafkaClient *kafka.KafkaInterface) *gin.Engine {

	//update with config env value
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	router.Use(gin.Recovery())

	api.Router(router, kafkaClient)

	return router
}
