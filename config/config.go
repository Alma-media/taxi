// Package config contains service settings
package config

import "time"

// Config contains settings of the service
type Config struct {
	HTTP HTTP

	Generator Generator

	Storage Storage
}

// HTTP server config
type HTTP struct {
	// Address to use for the API
	Address string
}

// Storage config
type Storage struct {
	// MaxSize contains the max number of elements (orders) to be stored
	MaxSize int
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
		HTTP: HTTP{
			Address: ":8080",
		},
		Storage: Storage{
			MaxSize: 1<<7 - 1,
		},
		Generator: Generator{
			KeySize:         2,
			KeyBytes:        "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			PoolSize:        50,
			ReplaceInterval: 200 * time.Millisecond,
		},
	}
}
