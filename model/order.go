package model

import (
	"fmt"
	"sync/atomic"
)

// Order model (could be a composite type but it was not required)
type Order struct {
	ID      string
	counter uint64
}

// NewOrder creates a new order with provided ID/key
func NewOrder(id string) *Order {
	return &Order{
		ID: id,
	}
}

func (o Order) String() string {
	return fmt.Sprintf("%s - %d", o.ID, atomic.LoadUint64(&o.counter))
}

// Increment order counter
func (o *Order) Increment() {
	atomic.AddUint64(&o.counter, 1)
}
