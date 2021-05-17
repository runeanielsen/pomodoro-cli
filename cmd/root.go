package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pomodoro-cli",
	Short: "Pomodoro cli is a command line tool for the pomodoro technique.",
	Long:  "Pomodoro cli is a command line tool to manage your workflow using the pomodoro technique.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
