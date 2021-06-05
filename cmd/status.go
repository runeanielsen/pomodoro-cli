package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Displays the status of the pomodoro",
	RunE: func(cmd *cobra.Command, args []string) error {
		return statusAction(os.Stdout)
	},
}

func statusAction(out io.Writer) error {
	fileName := "/tmp/pomodoro.json"

	p, err := pomodoro.LoadLatest(fileName)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	if p.Cancelled || p.HasEnded(now) {
		fmt.Fprint(out, "00:00\n")
	} else {
		fmt.Fprintf(out, "%s\n", pomodoro.FmtDuration(p.TimeLeft(now)))
	}

	return nil
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
