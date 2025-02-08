package order

import (
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {
	service := NewOrderService()
	handler := NewOrderHandler(service)

	g.POST("/", handler.HandlePostOrder)
}
