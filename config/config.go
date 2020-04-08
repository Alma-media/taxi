// Package config contains service settings
package config

import "time"

// Config contains settings of the service
type Config struct {
	HTTP HTTP

	Generator Generator
}

// HTTP server config
type HTTP struct {
	// Address to use for the API
	Address string `default:":8080"`
}

// Generator config
type Generator struct {
	// KeySize is a size of unique order identifier
	KeySize int `default:"2"`

	// KeyBytes contains allowed symbols to be used for key generation
	KeyBytes string `default:"ABCDEFGHIJKLMNOPQRSTUVWXYZ"`

	// PoolSize contains the length of order pool
	PoolSize int `default:"50"`

	// ReplaceInterval is an interval to replace an order in the pool with a new one
	ReplaceInterval time.Duration `default:"200ms"`
}
