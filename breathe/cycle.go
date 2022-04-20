package breathe

import (
	"fmt"
	"time"
)

// A BreatheCycle represents full cycle of breathing consisting of
// an inhale and an exhale
type BreatheCycle struct {
	Inhale     time.Duration
	InhaleHold time.Duration
	Exhale     time.Duration
	ExhaleHold time.Duration
}

// Display information about the breath cycles and run them
func RunBreatheCycles(cycle BreatheCycle, cyclesCount int) {
	totalDuration := time.Duration(cyclesCount*int((cycle.Inhale+cycle.InhaleHold+cycle.Exhale+cycle.ExhaleHold).Milliseconds())) * time.Millisecond
	fmt.Printf("%d breathing cycles with a total duration of %s\n", cyclesCount, totalDuration)

	fmt.Println("Each cycle consists of:")
	fmt.Printf("%.1f seconds inhale\n", float64(cycle.Inhale.Milliseconds())/1000)
	if cycle.InhaleHold.Milliseconds() > 0 {
		fmt.Printf("%.1f seconds hold\n", float64(cycle.InhaleHold.Milliseconds())/1000)
	}
	fmt.Printf("%.1f seconds exhale\n", float64(cycle.Exhale.Milliseconds())/1000)
	if cycle.ExhaleHold.Milliseconds() > 0 {
		fmt.Printf("%.1f seconds hold\n", float64(cycle.ExhaleHold.Milliseconds())/1000)
	}

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
	if cycle.InhaleHold.Milliseconds() > 0 {
		runBreatheSubCycle("Hold", cycle.InhaleHold)
	}

	runBreatheSubCycle("Exhale", cycle.Exhale)
	if cycle.ExhaleHold.Milliseconds() > 0 {
		runBreatheSubCycle("Hold", cycle.ExhaleHold)
	}
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
