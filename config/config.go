// Package config contains service settings
package config

import (
	"time"
)

// Config contains settings of the service
type Config struct {
	Generator Generator
}

// Generator config
type Generator struct {
	// KeySize is a size of unique order identifier
	KeySize int

	// KeyBytes contains allowed symbols to be used for key generation
	KeyBytes string

	// PoolSize contains the length of order pool
	PoolSize int

	// ReplaceInterval is an interval to replace an order in the pool with a new one
	ReplaceInterval time.Duration
}

// New creates a new config
// Further improvements:
// - parse the values from flags/env/yaml/toml/json/...
// - support default values
func New() Config {
	return Config{
		Generator: Generator{
			KeySize:         2,
			KeyBytes:        "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			PoolSize:        50,
			ReplaceInterval: 200 * time.Millisecond,
		},
	}
}
