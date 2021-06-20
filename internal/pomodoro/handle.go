package pomodoro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// The Pomodoro structure
type Pomodoro struct {
	Started   time.Time
	Duration  time.Duration
	Cancelled bool
}

// Calculates the end time.Time of the pomodoro based on the Started field
// and the DurationMins
func (p Pomodoro) EndTime() time.Time {
	return p.Started.Add(p.Duration)
}

// Get the time that is left of the pomodoro
func (p Pomodoro) TimeLeft(currentTime time.Time) time.Duration {
	left := p.EndTime().UTC().Sub(currentTime)
	if left > 0 {
		return left
	} else {
		return time.Second * 0
	}
}

// Return true if the pomodoro has ended
func (p Pomodoro) HasEnded(now time.Time) bool {
	return p.EndTime().UTC().Before(now.UTC())
}

// Sets the pomodoro
func (p *Pomodoro) Cancel(now time.Time) error {
	if p.Cancelled {
		return fmt.Errorf("The pomodoro is already cancelled.")
	}

	if p.HasEnded(now) {
		return fmt.Errorf("The pomodoro has ended and can not be cancelled.")
	}

	p.Cancelled = true

	return nil
}

// Creates a new pomodoro and adds is to the pomodoro list
func Add(fileName string, startTime time.Time, duration time.Duration) (Pomodoro, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		ioutil.WriteFile(fileName, nil, 0644)
	}

	l, err := LoadLatest(fileName)
	if err != nil {
		return Pomodoro{}, err
	}

	if !l.HasEnded(startTime) && !l.Cancelled {
		return Pomodoro{},
			fmt.Errorf("Cannot add new pomodoro, please cancel the current one or wait till it is completed.")
	}

	pomodoros, err := Load(fileName)
	if err != nil {
		return Pomodoro{}, nil
	}

	newPomodoro := Pomodoro{
		Started:  startTime,
		Duration: duration,
	}

	pomodoros = append(pomodoros, newPomodoro)

	err = Save(pomodoros, fileName)
	if err != nil {
		return Pomodoro{}, nil
	}

	return newPomodoro, nil
}

// Saves the list of pomodoros to the file (overwrites)
func Save(pomodoros []Pomodoro, fileName string) error {
	byteValue, err := json.Marshal(pomodoros)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, byteValue, 0644)
}

// Loads the list of pomodoro from the specified file
// and returns them as a slice
func Load(fileName string) ([]Pomodoro, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return []Pomodoro{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []Pomodoro{}, err
	}

	var pomodoros []Pomodoro
	if len(byteValue) != 0 {
		err = json.Unmarshal(byteValue, &pomodoros)
		if err != nil {
			return []Pomodoro{}, err
		}
	}

	return pomodoros, nil
}

// Loads the latest pomodoro
func LoadLatest(fileName string) (Pomodoro, error) {
	pomodoros, err := Load(fileName)
	if err != nil {
		return Pomodoro{}, nil
	}

	if len(pomodoros) == 0 {
		return Pomodoro{}, nil
	}

	return pomodoros[len(pomodoros)-1], nil
}
