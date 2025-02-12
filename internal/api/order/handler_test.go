package order

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-kafka-order-producer/internal/infra/events/kafka"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandlePostOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	kafkaMock := kafka.KafkaSpy{}
	Router(&gin.Default().RouterGroup, kafkaMock)
	path := "/order/v1/"

	t.Run("Should return error when payload is empty", func(t *testing.T) {
		orderService := NewOrderService(kafkaMock)
		orderHandler := NewOrderHandler(orderService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, nil)
		orderHandler.HandlePostOrder(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Should return error when the given params are of different types than the expected", func(t *testing.T) {
		mockRequest := map[string]interface{}{
			"client_id": 123,
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		orderService := NewOrderService(kafkaMock)
		orderHandler := NewOrderHandler(orderService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		orderHandler.HandlePostOrder(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, string(response), "cannot unmarshal")
	})

	t.Run("Should return a validation error", func(t *testing.T) {
		mockRequest := PostOrderRequest{
			Amount:   -123.00,
			ClientID: uuid.NewString(),
			StoreID:  uuid.NewString(),
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		orderService := NewOrderService(kafkaMock)
		orderHandler := NewOrderHandler(orderService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		orderHandler.HandlePostOrder(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, string(response), "amount must be a positive number")
	})

	t.Run("Should return an service error", func(t *testing.T) {
		errorMessage := "error to save order transaction"
		mockRequest := PostOrderRequest{
			Amount:            123.00,
			ClientID:          uuid.NewString(),
			StoreID:           uuid.NewString(),
			NotificationEmail: gofakeit.Email(),
		}

		orderService := OrderServiceSpy{
			PostOrderResponse: PostOrderResponse{},
			PostOrderError:    errors.New(errorMessage),
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		orderHandler := NewOrderHandler(orderService)

		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		orderHandler.HandlePostOrder(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, string(response), errorMessage)
	})

	t.Run("Should create the order", func(t *testing.T) {
		mockRequest := PostOrderRequest{
			Amount:            123.00,
			ClientID:          uuid.NewString(),
			StoreID:           uuid.NewString(),
			NotificationEmail: gofakeit.Email(),
		}

		orderService := OrderServiceSpy{
			PostOrderResponse: PostOrderResponse{
				uuid.NewString(),
			},
			PostOrderError: nil,
		}

		requestBytes, _ := json.Marshal(mockRequest)
		ioReader := bytes.NewBuffer(requestBytes)
		ioRequest := io.NopCloser(ioReader)

		orderHandler := NewOrderHandler(orderService)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, ioRequest)

		orderHandler.HandlePostOrder(ctx)

		response, _ := io.ReadAll(w.Body)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, string(response), "order_id")
	})
}
