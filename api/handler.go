// Package api provides HTTP handlers
// Further improvements:
// - enable the access control for public and private endpoints (use middleware)
// - retrieve codec from context to use encoder (e.g. codec.NewEncoder(w).Encode(order))
// - pass logger to the handler (intermal errors should be logged but not exposed to the user)
// - find approximate limit to enable trottling
package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Alma-media/taxi/model"
)

// OrderRepository represents a repository suitable for keeping orders
type OrderRepository interface {
	// Save should return random order from the abstract srotage
	Order(ctx context.Context) (*model.Order, error)

	// List should return a list of orders (generate a report)
	List(ctx context.Context) ([]*model.Order, error)
}

// NewHandler creates an HTTP handler
func NewHandler(repo OrderRepository) http.Handler {
	// since we do not need support for dynamic routes and parameters default ServeMux
	// is the most efficient choise (no need to use gin, gorilla, julienschmidt
	// or any other router based on RADIX tree)
	mux := http.NewServeMux()
	mux.Handle("/request/", CreateRequestHandler(repo))
	mux.Handle("/admin/requests", CreateAdminHandler(repo))

	return mux
}

// CreateRequestHandler creates a handler that returns a single random order
func CreateRequestHandler(repo OrderRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		order, err := repo.Order(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s\n", order.ID)
	}
}

// CreateAdminHandler creates a handler that generates a report
func CreateAdminHandler(repo OrderRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := repo.List(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, order := range list {
			fmt.Fprintf(w, "%s\n", order)
		}
	}
}
