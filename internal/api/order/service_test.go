package order

import (
	"errors"
	"go-kafka-order-producer/internal/infra/events/kafka"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	t.Run("should return an uuid when created a new order", func(t *testing.T) {
		orderService := NewOrderService(kafka.KafkaSpy{})
		response, err := orderService.PostOrder(ctx, &PostOrderRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
			StoreID:  uuid.NewString(),
		})

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Nil(t, uuid.Validate(response.OrderID))
	})

	t.Run("should return an error when calling kafka producer to create order", func(t *testing.T) {
		orderService := NewOrderService(kafka.KafkaSpy{
			ProduceError: errors.New("error to conect to kafka"),
		})
		response, err := orderService.PostOrder(ctx, &PostOrderRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
			StoreID:  uuid.NewString(),
		})

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error to conect to kafka")
	})

}
