package repository

import (
	"context"

	"github.com/tiny-go/taxi/model"
)

// Order represents a repository suitable for keeping orders
type Order interface {
	// Save should create a new order with provided key (if not exists) and save to
	// the storage or update an existing one returning the created/updated entity
	// Note that for some of the implementations context might not be required
	Save(ctx context.Context, req string) (*model.Order, error)

	// List should return a list of orders (generate a report)
	List(ctx context.Context) ([]*model.Order, error)
}
