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

// Pipe is a wrapper that combines orders source and storage
type Pipe struct {
	storage storage.Order
	source  Source
}

// NewPipe creates a "proxy" for provided generator in order to count
// the calls and store orders to the repository
func NewPipe(source Source, storage storage.Order) *Pipe {
	return &Pipe{
		storage: storage,
		source:  source,
	}
}

// Order wraps original Order() method
func (p *Pipe) Order(ctx context.Context) (*model.Order, error) {
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
func (p *Pipe) List(ctx context.Context) ([]*model.Order, error) {
	return p.storage.List(ctx)
}
