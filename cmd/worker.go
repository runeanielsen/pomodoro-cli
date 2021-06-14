package cmd

import (
	"time"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var workerCmd = &cobra.Command{
	Use:    "worker",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		pFile := viper.GetString("storage")
		fFile := viper.GetString("finished")
		return initWorker(pFile, fFile)
	},
}

func initWorker(pFile string, fFile string) error {
	startedPomo, err := pomodoro.LoadLatest(pFile)
	if err != nil {
		return err
	}

	time.Sleep(startedPomo.Duration)

	// We load it again to make sure it has not been canceled since it was started
	reloadedPomo, err := pomodoro.LoadLatest(pFile)
	if err != nil {
		return err
	}

	// If it has been canceled or it was canceled and started new one
	// We do this because the pomodoros does not have an id
	if reloadedPomo.Cancelled || startedPomo.Started != reloadedPomo.Started {
		return nil
	}

	if err = pomodoro.ExecuteHook(fFile); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
