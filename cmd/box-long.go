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

// boxLongCmd represents the box command
var boxLongCmd = &cobra.Command{
	Use:   "box-long",
	Short: "The long box breathing cycle",
	Long: `Navy SEALs use the box breathing cycle to calm themselves down.

It consists of 4 seconds inhale and 6 seconds exhale and 2 seconds holding
your breath in between.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting long box breathing cycle")
		fmt.Println("")

		breathe.RunBreatheCycles(
			breathe.BreatheCycle{
				Inhale:     4000 * time.Millisecond,
				InhaleHold: 4000 * time.Millisecond,
				Exhale:     6000 * time.Millisecond,
				ExhaleHold: 2000 * time.Millisecond,
			},
			20,
		)
	},
}

func init() {
	rootCmd.AddCommand(boxLongCmd)
}
