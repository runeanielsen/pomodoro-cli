package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		pFile := viper.GetString("storage")
		if err != nil {
			return err
		}

		return startAction(os.Stdout, duration, pFile)
	},
}

func startAction(out io.Writer, dMins int8, pFile string) error {
	now := time.Now().UTC()

	p, err := pomodoro.Start(pFile, now, dMins)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "Started pomodoro %s. The pomodoro will end %s.\n",
		p.Started.Local().Format("2 Jan 2006 15:04"),
		p.EndTime().Local().Format("2 Jan 2006 15:04"))

	bg := exec.Command("pomodoro-cli", "worker")
	if err = bg.Start(); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().Int8P("duration", "d", 25, "Duration of the pomodoro")
}
