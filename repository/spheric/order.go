package spheric

import (
	"context"
	"sort"
	"sync"

	"github.com/Alma-media/taxi/model"
	"github.com/Alma-media/taxi/repository"
)

// OrderRepository is an abstract implementation of order repository
type OrderRepository struct {
	mu     sync.RWMutex
	orders map[string]*model.Order
}

// NewOrderRepository creates a new order repository based on a map
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*model.Order),
	}
}

// Save the order (emulates "upsert" behavior)
// Always returns nil, context is not used in the current implementation but might be
// required by another driver, e.g. database)
func (repo *OrderRepository) Save(_ context.Context, key string) (*model.Order, error) {
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
func (repo *OrderRepository) List(_ context.Context) ([]*model.Order, error) {
	repo.mu.RLock()
	list := make([]*model.Order, 0, len(repo.orders))
	for _, order := range repo.orders {
		list = append(list, order)
	}
	repo.mu.RUnlock()

	sort.Sort(repository.ByID(list))

	return list, nil
}
