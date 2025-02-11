package order

import (
	"context"
	"go-kafka-order-producer/internal/infra/events/kafka"
)

type OrderServiceInterface interface {
	PostOrder(ctx context.Context, order *PostOrderRequest) (*PostOrderResponse, error)
}

type OrderService struct {
	KafkaProducer kafka.KafkaInterface
}

func NewOrderService(kafkaProducer kafka.KafkaInterface) *OrderService {
	return &OrderService{
		KafkaProducer: kafkaProducer,
	}
}

func (s *OrderService) PostOrder(ctx context.Context, order *PostOrderRequest) (*PostOrderResponse, error) {
	if err := s.KafkaProducer.Produce(ctx, "orders", "order.create", order); err != nil {
		return nil, err
	}

	return &PostOrderResponse{
		OrderID: order.OrderID,
	}, nil
}
