package repository

import (
	"context"

	"github.com/Alma-media/taxi/model"
	"github.com/Alma-media/taxi/storage"
)

// Source is an abstract source of orders
type Source interface {
	Generate() string
}

// OrderRepository combines orders source and storage
type OrderRepository struct {
	storage storage.Order
	source  Source
}

// NewOrderRepository creates a "proxy" for provided generator in order to count
// the calls and store orders to the repository
// Further improvements:
// - use event driven approach instead of global storage
func NewOrderRepository(source Source, storage storage.Order) *OrderRepository {
	return &OrderRepository{
		storage: storage,
		source:  source,
	}
}

// Order returns a random order
func (p *OrderRepository) Order(ctx context.Context) (*model.Order, error) {
	order, err := p.storage.Save(ctx, p.source.Generate())
	if err != nil {
		return nil, err
	}

	order.Increment()

	return order, nil
}

// List simply calls storage List() method
func (p *OrderRepository) List(ctx context.Context) ([]*model.Order, error) {
	return p.storage.List(ctx)
}
