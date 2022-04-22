package breathe

import (
	"testing"
	"time"
)

func TestGenerateBreathCycles(t *testing.T) {
	cycle := BreathCycle{
		Inhale: 5500 * time.Millisecond,
		Exhale: 5500 * time.Millisecond,
	}

	cycles := GenerateBreathCycles(cycle, 3)

	if len(cycles) != 3 {
		t.Fatalf("Length of cycles slice is not 3 but %d", len(cycles))
	}

	if cycles[0].Inhale != 5500*time.Millisecond || cycles[0].Exhale != 5500*time.Millisecond {
		t.Fatalf("First BreathCycle is not correct")
	}
	if cycles[1].Inhale != 5500*time.Millisecond || cycles[1].Exhale != 5500*time.Millisecond {
		t.Fatalf("Second BreathCycle is not correct")
	}
	if cycles[2].Inhale != 5500*time.Millisecond || cycles[2].Exhale != 5500*time.Millisecond {
		t.Fatalf("Third BreathCycle is not correct")
	}
}
