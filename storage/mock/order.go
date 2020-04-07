package mock

import (
	"context"

	"github.com/Alma-media/taxi/model"
	"github.com/Alma-media/taxi/storage"
)

var (
	_ storage.Order = OrderSuccess{}
	_ storage.Order = OrderFailure{}
)

// OrderSuccess mock
type OrderSuccess struct{}

// Save order
func (m OrderSuccess) Save(_ context.Context, key string) (*model.Order, error) {
	return model.NewOrder(key), nil
}

// List orders
func (m OrderSuccess) List(context.Context) ([]*model.Order, error) {
	return []*model.Order{}, nil
}

// OrderFailure mock
type OrderFailure struct{ Err error }

// Save order
func (m OrderFailure) Save(context.Context, string) (*model.Order, error) { return nil, m.Err }

// List orders
func (m OrderFailure) List(context.Context) ([]*model.Order, error) { return nil, m.Err }
