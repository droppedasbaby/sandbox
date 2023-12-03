package generators

import "context"

// MaxChanSize is the maximum size of the channel used to generate ids.
const MaxChanSize = 1000

// Generator is an interface that all generators must implement.
type Generator[T comparable] interface {
	GetID() T
	IsStateful() bool
	SetState(s T)
	GenerateIDs(ctx context.Context)
}
