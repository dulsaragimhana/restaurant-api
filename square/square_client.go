package square

import (
	"errors"
)

type SquareClient struct {
	AccessToken string
}

func NewClient(token string) *SquareClient {
	return &SquareClient{AccessToken: token}
}

func (s *SquareClient) CreateOrder(req interface{}) (string, error) {
	// Implement Square order creation logic
	return "order_id", nil
}

func (s *SquareClient) SubmitPayment(orderID string, req interface{}) error {
	// Implement Square payment submission
	return errors.New("not implemented")
}
