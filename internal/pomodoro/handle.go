package pomodoro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// The pomodoro structure
type pomodoro struct {
	Started      time.Time
	DurationMins int8
	Cancelled    bool
}

// Calculates the end time.Time of the pomodoro based on the Started field
// and the DurationMins
func (p pomodoro) End() time.Time {
	return p.Started.Add(time.Minute * time.Duration(p.DurationMins))
}

func (p pomodoro) TimeLeft() time.Duration {
	return p.End().UTC().Sub(time.Now().UTC())
}

// Formats the duration in the following format mm:ss
func FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	return fmt.Sprintf("%02d:%02d", m, s)
}

// Return true if the pomodoro has ended
func HasEnded(p pomodoro, now time.Time) bool {
	return p.End().UTC().Before(now.UTC())
}

// Creates a new pomodoro and adds is to the pomodoro list
func Start(fileName string, startTime time.Time, dMins int8) (pomodoro, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		ioutil.WriteFile(fileName, nil, 0644)
	}

	l, err := LoadLatest(fileName)
	if err != nil {
		return pomodoro{}, err
	}

	if !HasEnded(l, startTime) || l.Cancelled {
		return pomodoro{},
			fmt.Errorf("Cannot start new pomodoro, please cancel the current one or wait till it is completed.")
	}

	pomodoros, err := Load(fileName)
	if err != nil {
		return pomodoro{}, nil
	}

	newPomodoro := pomodoro{
		Started:      startTime,
		DurationMins: dMins,
	}

	pomodoros = append(pomodoros, newPomodoro)

	err = save(pomodoros, fileName)
	if err != nil {
		return pomodoro{}, nil
	}

	return newPomodoro, nil
}

func save(pomodoros []pomodoro, fileName string) error {
	byteValue, err := json.Marshal(pomodoros)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, byteValue, 0644)
}

// Loads the list of pomodoro from the specified file
// and returns them as a slice
func Load(fileName string) ([]pomodoro, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return []pomodoro{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []pomodoro{}, err
	}

	var pomodoros []pomodoro
	if len(byteValue) != 0 {
		err = json.Unmarshal(byteValue, &pomodoros)
		if err != nil {
			return []pomodoro{}, err
		}
	}

	return pomodoros, nil
}

// Loads the latest pomodoro
func LoadLatest(fileName string) (pomodoro, error) {
	pomodoros, err := Load(fileName)
	if err != nil {
		return pomodoro{}, nil
	}

	if len(pomodoros) == 0 {
		return pomodoro{}, nil
	}

	return pomodoros[len(pomodoros)-1], nil
}
