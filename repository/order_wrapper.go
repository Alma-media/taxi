package repository

import (
	"context"

	"github.com/Alma-media/taxi/model"
)

// Source is an abstract source of orders
type Source interface {
	Generate(ctx context.Context) (string, error)
}

// OrderWrapper is a wrapper that combines orders source and storage
type OrderWrapper struct {
	OrderRepository

	source Source
}

// NewOrderWrapper creates a "proxy" for provided generator in order to count
// the calls and store orders to the repository
func NewOrderWrapper(source Source, original OrderRepository) *OrderWrapper {
	return &OrderWrapper{
		OrderRepository: original,
		source:          source,
	}
}

// Order wraps original Order() func
func (p *OrderWrapper) Order(ctx context.Context) (*model.Order, error) {
	key, err := p.source.Generate(ctx)
	if err != nil {
		return nil, err
	}

	order, err := p.OrderRepository.Save(ctx, key)
	if err != nil {
		return nil, err
	}

	order.Increment()

	return order, nil
}
