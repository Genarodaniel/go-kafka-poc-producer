package order

type OrderServiceSpy struct {
	OrderServiceInterface
	PostOrderResponse PostOrderResponse
	PostOrderError    error
}

func (s OrderServiceSpy) PostOrder(order *PostOrderRequest) (*PostOrderResponse, error) {
	return &s.PostOrderResponse, s.PostOrderError
}
