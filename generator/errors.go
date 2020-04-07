package generator

import "errors"

var (
	errNotInitialized = errors.New("storage was not initialized")

	errAlreadyInitialized = errors.New("storage was already initialized")
)
