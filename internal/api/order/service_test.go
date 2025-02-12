package order

import (
	"go-kafka-order-producer/internal/infra/events/kafka"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/twmb/franz-go/pkg/kgo"
)

func TestPostOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	t.Run("should return an uuid when created a new order", func(t *testing.T) {
		orderService := NewOrderService(kafka.NewKafka(&kgo.Client{}))
		response, err := orderService.PostOrder(ctx, &PostOrderRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
			StoreID:  uuid.NewString(),
		})

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Nil(t, uuid.Validate(response.OrderID))
	})

}
