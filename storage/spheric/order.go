package spheric

import (
	"context"
	"sort"
	"sync"

	"github.com/Alma-media/taxi/model"
	"github.com/Alma-media/taxi/repository"
)

// OrderStorage is an abstract storage implementation
type OrderStorage struct {
	mu     sync.RWMutex
	orders map[string]*model.Order
}

// NewOrderStorage creates a new order repository based on a map
func NewOrderStorage() *OrderStorage {
	return &OrderStorage{
		orders: make(map[string]*model.Order),
	}
}

// Save the order (emulates "upsert" behavior)
// Always returns nil, context is not used in the current implementation but might be
// required by another driver, e.g. database)
func (repo *OrderStorage) Save(_ context.Context, key string) (*model.Order, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	order, ok := repo.orders[key]
	if ok {
		return order, nil
	}

	order = model.NewOrder(key)
	repo.orders[key] = order

	return order, nil
}

// List retrieves all stored entities sorted by ID (current implementation cannot
// fail retrieving the list)
func (repo *OrderStorage) List(_ context.Context) ([]*model.Order, error) {
	repo.mu.RLock()
	list := make([]*model.Order, 0, len(repo.orders))
	for _, order := range repo.orders {
		list = append(list, order)
	}
	repo.mu.RUnlock()

	sort.Sort(repository.ByID(list))

	return list, nil
}
