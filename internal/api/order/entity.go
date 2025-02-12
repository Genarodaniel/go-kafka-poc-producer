package order

import (
	"errors"
	"go-kafka-order-producer/internal/infra/utils"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusCreated OrderStatus = "created"
)

type PostOrderResponse struct {
	OrderID string `json:"order_id"`
}

type PostOrderRequest struct {
	ClientID          string  `json:"client_id"`
	StoreID           string  `json:"store_id"`
	NotificationEmail string  `json:"notification_email"`
	Status            string  `json:"status"`
	OrderID           string  `json:"order_id"`
	Amount            float64 `json:"amount"`
}

func (request *PostOrderRequest) Validate() error {
	if request.Amount <= 0 {
		return errors.New("amount must be a positive number")
	}

	if err := uuid.Validate(request.ClientID); err != nil {
		return errors.New("client_id must be a uuid")
	}

	if err := uuid.Validate(request.StoreID); err != nil {
		return errors.New("store_id must be a uuid")
	}

	if valid := utils.ValidateEmail(request.NotificationEmail); !valid {
		return errors.New("notification_email invalid")
	}

	return nil
}
