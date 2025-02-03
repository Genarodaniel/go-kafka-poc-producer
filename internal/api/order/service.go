package order

import "github.com/google/uuid"

type OrderServiceInterface interface {
	PostOrder(order *PostOrderRequest) (*PostOrderResponse, error)
}

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (a *OrderService) PostOrder(order *PostOrderRequest) (*PostOrderResponse, error) {
	return &PostOrderResponse{
		OrderID: uuid.NewString(),
	}, nil
}
