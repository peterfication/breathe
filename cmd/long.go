/*
Copyright © 2022 Peter Gundel <mail@petergundel.de>

*/
package cmd

import (
	"fmt"
	"time"

	"breathe/breathe"

	"github.com/spf13/cobra"
)

var (
	inhaleStartSeconds    int
	inhaleEndSeconds      int
	cyclesPerStartSeconds int
)

// longCmd represents the long command
var longCmd = &cobra.Command{
	Use:   "long",
	Short: "Train long breathing",
	Long:  `Start short and increase to longer breaths.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting long breathing cycle")
		fmt.Println("")

		cycles := generateCycles(inhaleStartSeconds, inhaleEndSeconds, cyclesPerStartSeconds)

		fmt.Printf("%d breathing cycles with a total duration of %s\n", len(cycles), breathe.TotalDuration(cycles))

		for i, cycle := range cycles {
			fmt.Printf("Cycle %d of %d\n", i, len(cycles))
			breathe.RunBreatheCycle(cycle)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(longCmd)

	longCmd.Flags().IntVarP(&inhaleStartSeconds, "inhaleStartSeconds", "s", 4, "The amount of seconds inhaling to start")
	longCmd.Flags().IntVarP(&inhaleEndSeconds, "inhaleEndSeconds", "e", 10, "The amount of seconds inhaling to increase up to")
	longCmd.Flags().IntVarP(&cyclesPerStartSeconds, "cyclesPerStartSeconds", "c", 3, "The cycles count per inhaling seconds")
}

// Generate the cycles according to the inhalation seconds and the cycles provided.
func generateCycles(inhaleStartSeconds, inhaleEndSeconds, cyclesPerStartSeconds int) []breathe.BreatheCycle {
	var cycles = []breathe.BreatheCycle{}

	for i := inhaleStartSeconds; i <= inhaleEndSeconds; i++ {
		for j := 1; j <= cyclesPerStartSeconds; j++ {
			cycles = append(cycles,
				breathe.BreatheCycle{
					Inhale: time.Duration(i*1000) * time.Millisecond,
					Exhale: time.Duration(i*2*1000) * time.Millisecond,
				},
			)
		}
	}

	return cycles
}
