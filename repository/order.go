package repository

import (
	"context"

	"github.com/Alma-media/taxi/model"
)

// OrderRepository represents a repository suitable for keeping orders
type OrderRepository interface {
	// Save should return random order from the abstract srotage
	Order(ctx context.Context) (*model.Order, error)

	// List should return a list of orders (generate a report)
	List(ctx context.Context) ([]*model.Order, error)
}
