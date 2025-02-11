package server

import (
	"go-kafka-order-producer/internal/api"

	"github.com/gin-gonic/gin"
	"github.com/twmb/franz-go/pkg/kgo"
)

func Init(kafkaClient *kgo.Client) *gin.Engine {

	//update with config env value
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.LoggerWithWriter(gin.DefaultWriter))
	router.Use(gin.Recovery())

	api.Router(router, kafkaClient)

	return router
}
