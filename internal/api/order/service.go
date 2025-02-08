package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type OrderServiceInterface interface {
	PostOrder(ctx context.Context, order *PostOrderRequest) (*PostOrderResponse, error)
}

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (a *OrderService) PostOrder(ctx context.Context, order *PostOrderRequest) (*PostOrderResponse, error) {
	fmt.Println(order)

	return &PostOrderResponse{
		OrderID: uuid.NewString(),
	}, nil
}
