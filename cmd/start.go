package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "Start pomomdoro",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := startAction(os.Stdout)
		if err != nil {
			return err
		}

		return nil
	},
}

func startAction(out io.Writer) error {
	fileName := "/tmp/pomodoro.json"
	p, err := pomodoro.Start(fileName)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "Started pomodoro %s. The pomodor will end %s.\n",
		p.Started.Local().Format("2 Jan 2006 15:04"),
		p.End().Local().Format("2 Jan 2006 15:04"))

	return nil
}

func init() {
	rootCmd.AddCommand(startCmd)
}
