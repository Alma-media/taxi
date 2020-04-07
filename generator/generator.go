package generator

import (
	"context"
	"time"

	"github.com/Alma-media/taxi/config"
)

// Generator is an abstract source of orders
type Generator struct {
	config       config.Generator
	queue        chan string
	stop, status chan struct{}
}

// New cretaes a new storage
func New(cfg config.Generator) *Generator {
	return &Generator{
		config: cfg,
		status: make(chan struct{}, 1),
	}
}

// Init the source of orders
func (g *Generator) Init() error {
	select {
	case g.status <- struct{}{}:
	default:
		return errAlreadyInitialized
	}

	g.queue = make(chan string, g.config.PoolSize)
	g.stop = make(chan struct{})

	go g.fill(g.config.KeySize, g.config.KeyBytes)
	go g.pull(g.config.ReplaceInterval)

	return nil
}

// Generate picks an order from the queue
func (g *Generator) Generate(ctx context.Context) (string, error) {
	select {
	case key := <-g.queue:
		return key, nil
	// we have to handle the case when the storage has not been initialized / unavailable
	// (simulate storage failure)
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// Close the source of orders
func (g *Generator) Close() error {
	select {
	case <-g.status:
	default:
		return errNotInitialized
	}

	close(g.stop)
	return nil
}

// ensures that the queue is always full (has N random orders)
func (g *Generator) fill(length int, from string) {
	for {
		select {
		// generate and put a new key to the queue once it gets an empty slot
		case g.queue <- RandKey(length, from):
		case <-g.stop:
			return
		}
	}
}

// pull (or cancer) an order by ticker
func (g *Generator) pull(interval time.Duration) {
	tick := time.NewTicker(interval)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			// pick the order that needs to be canceled (the new one will be automatically
			// added by fill() once we get an empty slot)
			<-g.queue
		case <-g.stop:
			return
		}
	}
}
