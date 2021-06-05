package pomodoro

import (
	"fmt"
	"time"
)

// Formats the duration in the following format mm:ss
func FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	return fmt.Sprintf("%02d:%02d", m, s)
}
