/*
Copyright © 2022 Peter Gundel <mail@petergundel.de>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// idealCmd represents the ideal command
var idealCmd = &cobra.Command{
	Use:   "ideal",
	Short: "The ideal way to breathe",
	Long: `According to studies, the ideal way to breathe is inhale for
5.5 seconds and exhale for 5.5 seconds.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting ideal breathing cycle")
		fmt.Println("")

		runBreatheCycles(
			BreatheCycle{
				Inhale: 5500 * time.Millisecond,
				Exhale: 5500 * time.Millisecond,
			},
			55,
		)
	},
}

func init() {
	rootCmd.AddCommand(idealCmd)
}

// A BreatheCycle represents full cycle of breathing consisting of
// an inhale and an exhale
type BreatheCycle struct {
	Inhale time.Duration
	Exhale time.Duration
}

// Display information about the breath cycles and run them
func runBreatheCycles(cycle BreatheCycle, cyclesCount int) {
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
