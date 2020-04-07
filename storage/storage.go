package storage

import (
	"context"

	"github.com/Alma-media/taxi/model"
)

// Order represents persistent order storage
type Order interface {
	Save(ctx context.Context, key string) (*model.Order, error)
	List(ctx context.Context) ([]*model.Order, error)
}
