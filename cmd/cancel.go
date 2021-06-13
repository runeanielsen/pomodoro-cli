package cmd

import (
	"time"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:          "cancel",
	Short:        "Cancel the pomodoro",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		pFile := viper.GetString("storage")
		return cancelAction(pFile)
	},
}

func cancelAction(pFile string) error {
	pomodoros, err := pomodoro.Load(pFile)
	if err != nil {
		return err
	}

	p := &pomodoros[len(pomodoros)-1]
	if err := p.Cancel(time.Now().UTC()); err != nil {
		return err
	}

	if err := pomodoro.Save(pomodoros, pFile); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(cancelCmd)
}
