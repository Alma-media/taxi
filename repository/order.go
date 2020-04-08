package repository

import (
	"context"

	"github.com/Alma-media/taxi/model"
	"github.com/Alma-media/taxi/storage"
)

// Source is an abstract source of orders
type Source interface {
	Generate(ctx context.Context) (string, error)
}

// OrderRepository combines orders source and storage
type OrderRepository struct {
	storage storage.Order
	source  Source
}

// NewOrderRepository creates a "proxy" for provided generator in order to count
// the calls and store orders to the repository
// Further improvements:
// - use event driven architecture instead of global storage
func NewOrderRepository(source Source, storage storage.Order) *OrderRepository {
	return &OrderRepository{
		storage: storage,
		source:  source,
	}
}

// Order wraps original Order() method
func (p *OrderRepository) Order(ctx context.Context) (*model.Order, error) {
	key, err := p.source.Generate(ctx)
	if err != nil {
		return nil, err
	}

	order, err := p.storage.Save(ctx, key)
	if err != nil {
		return nil, err
	}

	order.Increment()

	return order, nil
}

// List wrap original List() method
func (p *OrderRepository) List(ctx context.Context) ([]*model.Order, error) {
	return p.storage.List(ctx)
}
