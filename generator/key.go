package generator

import (
	"math/rand"
	"time"
)

// RandKey creates a key of fixed size with provided symbols
func RandKey(size int, from string) string {
	bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = from[rand.Intn(len(from)-1)]
	}
	return string(bytes)
}

func init() { rand.Seed(time.Now().UnixNano()) }
