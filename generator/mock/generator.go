package mock

import "context"

// GeneratorSuccess generates a sequence with success
type GeneratorSuccess chan string

// Generate order sequence
func (m GeneratorSuccess) Generate(context.Context) (string, error) { return <-m, nil }

// GeneratorFailure fails to generate the sequence
type GeneratorFailure struct{ Err error }

// Generate order sequence
func (m GeneratorFailure) Generate(context.Context) (string, error) { return "", m.Err }
