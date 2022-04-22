/*
Copyright © 2022 Peter Gundel <mail@petergundel.de>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "breathe",
	Short: "Guidance in breathing",
	Long: `A tool that helps you in breathing certain ways.

Always inhale through the nose!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var sound string

func init() {
	rootCmd.PersistentFlags().StringVar(&sound, "sound", "none", "Whether to play sound none|words|numbers|all")
}
