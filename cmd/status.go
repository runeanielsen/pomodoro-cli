package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Displays the status of the pomodoro",
	RunE: func(cmd *cobra.Command, args []string) error {
		pFile := viper.GetString("storage")
		return statusAction(os.Stdout, pFile)
	},
}

func statusAction(out io.Writer, pFile string) error {
	p, err := pomodoro.LoadLatest(pFile)
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
