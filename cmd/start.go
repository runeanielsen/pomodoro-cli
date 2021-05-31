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
		duration, err := cmd.Flags().GetInt8("duration")
		if err != nil {
			return err
		}
		return startAction(os.Stdout, duration)
	},
}

func startAction(out io.Writer, dMins int8) error {
	fileName := "/tmp/pomodoro.json"

	p, err := pomodoro.Start(fileName, time.Now().UTC(), dMins)
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
	startCmd.Flags().Int8P("duration", "d", 25, "Duration of the pomodoro")
}
