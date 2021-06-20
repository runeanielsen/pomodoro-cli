package pomodoro

import (
	"time"
)

type TimerEvent string

const (
	FinishedBreak    TimerEvent = "BREAK"
	FinishedPomodoro TimerEvent = "POMODORO"
)

func PomdoroLoop(fFile string, pTime time.Duration, bTime time.Duration) {
	timerCh := make(chan TimerEvent)

	go func() {
		timer(timerCh, pTime, FinishedPomodoro)
		for timerEvent := range timerCh {
			switch timerEvent {
			case FinishedBreak:
				go ExecuteHook(fFile)
				timer(timerCh, pTime, FinishedPomodoro)
			case FinishedPomodoro:
				go ExecuteHook(fFile)
				timer(timerCh, bTime, FinishedBreak)
			}
		}
	}()
}

func timer(ch chan TimerEvent, d time.Duration, finishedEvent TimerEvent) {
	time.AfterFunc(d, func() {
		ch <- finishedEvent
	})
}
