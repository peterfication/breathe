package breathe

import (
	"fmt"
	"time"
)

// A BreatheCycle represents full cycle of breathing consisting of
// an inhale and an exhale
type BreatheCycle struct {
	Inhale time.Duration
	Exhale time.Duration
}

// Display information about the breath cycles and run them
func RunBreatheCycles(cycle BreatheCycle, cyclesCount int) {
	fmt.Printf(
		"%d cycles with %.1f seconds inhale followed by %.1f seconds exhale.\n",
		cyclesCount,
		float64(cycle.Inhale.Milliseconds())/1000,
		float64(cycle.Exhale.Milliseconds())/1000,
	)
	totalDuration := time.Duration(cyclesCount*int((cycle.Inhale+cycle.Exhale).Milliseconds())) * time.Millisecond
	fmt.Printf("Total duration: %s\n", totalDuration)
	fmt.Println("")

	time.Sleep(1 * time.Second)

	for i := 0; i < cyclesCount; i++ {
		fmt.Printf("Cycle %d of %d\n", i, cyclesCount)
		runBreatheCycle(cycle)
		fmt.Println()
	}
}

// Run a single breath cycle consisting of an inhale and an exhale step
func runBreatheCycle(cycle BreatheCycle) {
	runBreatheSubCycle("Inhale", cycle.Inhale)
	runBreatheSubCycle("Exhale", cycle.Exhale)
}

// Run a single breath sub cycle like an inhale or an exhale step
// by waiting the appropriate time and printing information about
// how long is still to go.
func runBreatheSubCycle(startWord string, duration time.Duration) {
	fmt.Printf("%s ", startWord)
	for i := int(duration.Milliseconds() / 100); i > 0; i-- {
		time.Sleep(100 * time.Millisecond)

		firstSecond := int(duration.Milliseconds()/100) - 10 - int(duration.Milliseconds()/100)%10 + 1
		if i < firstSecond && i%10 == 0 {
			fmt.Printf("%d ", i/10)
		}
	}
	fmt.Println()
}
