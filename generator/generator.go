package generator

import (
	"context"
	"time"

	"github.com/Alma-media/taxi/config"
)

// Generator is an abstract source of orders
type Generator struct {
	config config.Generator
	queue  chan string
}

// New cretaes a new storage
func New(cfg config.Generator) *Generator {
	return &Generator{
		config: cfg,
		queue:  make(chan string),
	}
}

// Run the generator
func (g *Generator) Run(ctx context.Context) {
	go g.fill(ctx)
	go g.pull(ctx)
}

// Generate picks an order from the queue
func (g *Generator) Generate() string {
	return <-g.queue
}

// ensures that the queue is always full (has N random orders)
func (g *Generator) fill(ctx context.Context) {
	for {
		select {
		// generate and put a new key to the queue once it gets an empty slot
		case g.queue <- RandKey(g.config.KeySize, g.config.KeyBytes):
		case <-ctx.Done():
			return
		}
	}
}

// pull (or cancer) an order by ticker
func (g *Generator) pull(ctx context.Context) {
	tick := time.NewTicker(g.config.ReplaceInterval)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			// pick the order that needs to be canceled (the new one will be automatically
			// added by fill() once we get an empty slot)
			<-g.queue
		case <-ctx.Done():
			return
		}
	}
}
