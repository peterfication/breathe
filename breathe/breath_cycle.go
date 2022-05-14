package breathe

import (
	"time"
)

// A BreathCycle represents full cycle of breathing consisting of
// an inhale and an exhale
type BreathCycle struct {
	Inhale     time.Duration
	InhaleHold time.Duration
	Exhale     time.Duration
	ExhaleHold time.Duration
}

// Take one BreathCycle and create a slice of the same BreathCycles with
// the cyclesCount of it.
func GenerateBreathCycles(breathCycle BreathCycle, cyclesCount int) (breathCycles []BreathCycle) {
	for i := 0; i < cyclesCount; i++ {
		breathCycles = append(breathCycles, breathCycle)
	}
	return breathCycles
}
