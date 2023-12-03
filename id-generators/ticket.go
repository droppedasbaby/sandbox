package generators

import (
	"context"
	"math"
)

// TicketGenerator is a generator that generates a new ticket id.
// This is somewhat based off of Flickr's ticket system.
// https://code.flickr.net/2010/02/08/ticket-servers-distributed-unique-primary-keys-on-the-cheap/
type TicketGenerator struct {
	curr uint64
	ch   chan uint64
}

// NewTicketGenerator creates a new TicketGenerator.
func NewTicketGenerator() *TicketGenerator {
	return &TicketGenerator{
		curr: 1,
		ch:   make(chan uint64, MaxChanSize),
	}
}

// NewTicketGeneratorFromInt creates a new TicketGenerator.
// The starting value is the value of i. To be used when restarting the generator.
func NewTicketGeneratorFromInt(i uint64) *TicketGenerator {
	return &TicketGenerator{
		curr: i,
		ch:   make(chan uint64, MaxChanSize),
	}
}

// GenerateIDs is a long-running method that generates ticket ids and sends them to the channel.
// Is stopped when the context is cancelled.
func (tg *TicketGenerator) GenerateIDs(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			tg.ch <- tg.curr
			if tg.curr == math.MaxUint64 {
				panic("TicketGenerator.Curr will overflow if incremented.")
			}
			tg.curr++
		}
	}
}

// GetID returns the next ticket id.
func (tg *TicketGenerator) GetID() uint64 {
	return <-tg.ch
}

// IsStateful returns true since this generator is stateful, ie the ids are sequential ints.
func (tg *TicketGenerator) IsStateful() bool {
	return true
}

// SetState sets the state of the generator, which is the current ticket id.
func (tg *TicketGenerator) SetState(s uint64) {
	tg.curr = s
}
