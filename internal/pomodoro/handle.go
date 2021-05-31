package pomodoro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

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

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	return fmt.Sprintf("%02d:%02d", m, s)
}

func HasEnded(p pomodoro, now time.Time) bool {
	return p.End().UTC().Before(now.UTC())
}

func Start(fileName string, startTime time.Time) (pomodoro, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		ioutil.WriteFile(fileName, nil, 0644)
	}

	pomodoros, err := Load(fileName)
	if err != nil {
		return pomodoro{}, nil
	}

	newPomodoro := pomodoro{
		Started:      startTime,
		DurationMins: 25,
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

// Loads the latest pomodoro,
// if no pomodoro is available returns error that the list is empty
func LoadLatest(fileName string) (pomodoro, error) {
	pomodoros, err := Load(fileName)
	if err != nil {
		return pomodoro{}, nil
	}

	if len(pomodoros) == 0 {
		return pomodoro{}, fmt.Errorf("The list of pomodoros is empty.")
	}

	return pomodoros[len(pomodoros)-1], nil
}
