/*
Copyright © 2022 Peter Gundel <mail@petergundel.de>

*/
package cmd

import (
	"time"

	"breathe/breathe"

	"github.com/spf13/cobra"
)

// boxCmd represents the box command
var boxCmd = &cobra.Command{
	Use:   "box",
	Short: "The box breathing cycle",
	Long: `Navy SEALs use the box breathing cycle to calm themselves down.

It consists of 4 seconds inhale and 4 seconds exhale and 4 seconds holding
your breath in between.`,
	Run: func(cmd *cobra.Command, args []string) {
		breathe.RunBreathCycles(
			"Box breathe cycle (4x4s)",
			breathe.GenerateBreathCycles(
				breathe.BreathCycle{
					Inhale:     4000 * time.Millisecond,
					InhaleHold: 4000 * time.Millisecond,
					Exhale:     4000 * time.Millisecond,
					ExhaleHold: 4000 * time.Millisecond,
				},
				20,
			),
			rootCmd.PersistentFlags().Lookup("sound").Value.String(),
		)
	},
}

func init() {
	rootCmd.AddCommand(boxCmd)
}
