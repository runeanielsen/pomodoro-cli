package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "Start pomomdoro",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startAction(os.Stdout)
	},
}

func startAction(out io.Writer) error {
	fileName := "/tmp/pomodoro.json"

	p, err := pomodoro.Start(fileName, time.Now().UTC())
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
