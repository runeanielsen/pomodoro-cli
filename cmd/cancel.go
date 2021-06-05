package cmd

import (
	"time"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:          "cancel",
	Short:        "Cancel the pomodoro",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cancelAction()
	},
}

func cancelAction() error {
	fileName := "/tmp/pomodoro.json"
	pomodoros, err := pomodoro.Load(fileName)
	if err != nil {
		return err
	}

	p := &pomodoros[len(pomodoros)-1]
	if err := p.Cancel(time.Now().UTC()); err != nil {
		return err
	}

	if err := pomodoro.Save(pomodoros, fileName); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(cancelCmd)
}
