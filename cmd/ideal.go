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

// idealCmd represents the ideal command
var idealCmd = &cobra.Command{
	Use:   "ideal",
	Short: "The ideal way to breathe",
	Long: `According to studies, the ideal way to breathe is inhale for
5.5 seconds and exhale for 5.5 seconds.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting ideal breathing cycle")
		fmt.Println("")

		breathe.RunBreatheCycles(
			breathe.BreatheCycle{
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
