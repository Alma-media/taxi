package mock

// Generator generates a sequence of orders
type Generator chan string

// Generate order sequence
func (m Generator) Generate() string { return <-m }
