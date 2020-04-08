package mock

import (
	"context"

	"github.com/Alma-media/taxi/model"
)

// OrderRepository mock
type OrderRepository struct {
	Err           error
	OrderResponse *model.Order
	ListResponse  []*model.Order
}

// Order returns a single order or fails with provided error
func (m OrderRepository) Order(context.Context) (*model.Order, error) {
	return m.OrderResponse, m.Err
}

// List returns the list of orders with statistics or fails with provided error
func (m OrderRepository) List(context.Context) ([]*model.Order, error) {
	return m.ListResponse, m.Err
}
