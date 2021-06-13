package pomodoro

import (
	"os"
	"os/exec"
)

// Runs the hook if the file exists
func ExecuteHook(hookPath string) error {
	if _, err := os.Stat(hookPath); !os.IsNotExist(err) {
		cmd := exec.Command("/bin/sh", hookPath)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
