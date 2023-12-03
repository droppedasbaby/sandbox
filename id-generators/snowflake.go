package generators

import (
	"context"
	"sync"
	"time"
)

// SnowflakeGenerator is a generator that generates snowflake ids.
// Returns 64-bit unsigned ints.
// This is somewhat based off of Twitter's snowflake system.
// https://en.wikipedia.org/wiki/Snowflake_ID
type SnowflakeGenerator struct {
	epoch             uint64
	timestampBits     uint8
	machineBits       uint8
	machineID         uint64
	sequenceBits      uint8
	mutex             sync.Mutex
	lastTimestamp     uint64
	sequence          uint64
	sequenceResetNano uint64
}

const MilliToNano = 1000000

// NewSnowflakeGenerator creates a new SnowflakeGenerator.
// Each part of the snowflake is configurable, but the default values are the same as Twitter's.
// To configure, use NewSnowflakeGeneratorFromConfig.
func NewSnowflakeGenerator() *SnowflakeGenerator {
	machineBits := uint8(10)
	return &SnowflakeGenerator{
		epoch:             0,
		timestampBits:     41,
		machineBits:       machineBits,
		machineID:         GenerateMachineID(machineBits),
		sequenceBits:      12,
		sequenceResetNano: MilliToNano,
	}
}

// NewSnowflakeGeneratorFromConfig creates a new SnowflakeGenerator.
// This is the same as NewSnowflakeGenerator, but the bits are configurable.
func NewSnowflakeGeneratorFromConfig(
	epoch uint64, timestampBits uint8, machineBits uint8, machineID uint64, sequenceBits uint8,
) *SnowflakeGenerator {
	return &SnowflakeGenerator{
		epoch:         epoch,
		timestampBits: timestampBits,
		machineBits:   machineBits,
		machineID:     machineID,
		sequenceBits:  sequenceBits,
	}
}

// GetID returns a snowflake id.
func (sg *SnowflakeGenerator) GetID() uint64 {
	sg.mutex.Lock()
	defer sg.mutex.Unlock()

	now := time.Now()
	timestampNano := uint64(now.UnixNano()) - sg.epoch
	if timestampNano-sg.lastTimestamp >= sg.sequenceResetNano {
		sg.sequence = 0
	} else {
		sg.sequence++
		if sg.sequence >= uint64(1<<sg.sequenceBits) {
			panic("Sequence overflow")
		}
	}

	sg.lastTimestamp = timestampNano

	id := timestampNano / MilliToNano
	id <<= sg.machineBits
	id |= sg.machineID
	id <<= sg.sequenceBits
	id |= sg.sequence

	return id
}

// IsStateful returns false since this generator is not stateful, ie the ids are not sequential ints.
func (sg *SnowflakeGenerator) IsStateful() bool {
	return false
}

// GenerateIDs is a long-running method that generates snowflake ids and sends them to the channel.
// Should be called if the generator is stateful.
func (sg *SnowflakeGenerator) GenerateIDs(ctx context.Context) {
	panic("SnowflakeGenerator.GenerateIDs should not be called.")
}

// SetState sets the state of the generator, panics since this generator is not stateful.
func (sg *SnowflakeGenerator) SetState(s uint64) {
	panic("SnowflakeGenerator.SetState should not be called.")
}
