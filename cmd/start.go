package cmd

import (
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type StartParameters struct {
	PomodoroDuration time.Duration
	PomodoroFilePath string
}

var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "Start pomomdoro",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		mins, err := cmd.Flags().GetInt8("duration")
		if err != nil {
			return err
		}

		silent, err := cmd.Flags().GetBool("silent")
		if err != nil {
			return err
		}

		fFile := viper.GetString("finished")
		pFile := viper.GetString("storage")

		return startAction(os.Stdout, mins, pFile, fFile, silent)
	},
}

func startAction(out io.Writer, mins int8, pFile string, fFile string, silent bool) error {
	pomodoro.PomdoroLoop(fFile, time.Duration(mins)*time.Minute, time.Duration(5)*time.Minute)

	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	sig := <-signalCh
	switch sig {
	case os.Interrupt:
		return nil
	case syscall.SIGTERM:
		return nil
	}

	return nil
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().Int8P("pduration", "d", 25, "Duration of the pomodoros")
	startCmd.Flags().Int8P("bduration", "b", 5, "Duration of the breaks")
	startCmd.Flags().BoolP("silent", "s", false, "Silence the output")
}
