package breathe

import (
	"testing"
)

func TestBreathCyclesCount(t *testing.T) {
	runner := Runner{}

	runner.breathCycles = []BreathCycle{BreathCycle{}, BreathCycle{}}
	if runner.BreathCyclesCount() != 2 {
		t.Fatalf("breath cycles count should be 2")
	}

	runner.breathCycles = []BreathCycle{}
	if runner.BreathCyclesCount() != 0 {
		t.Fatalf("breath cycles count should be 0")
	}
}
