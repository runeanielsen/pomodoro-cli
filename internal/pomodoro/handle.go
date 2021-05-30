package pomodoro

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type pomodoro struct {
	Started      time.Time
	DurationMins int8
}

// Calculates the end time.Time of the pomodoro based on the Started field
// and the DurationMins
func (p pomodoro) End() time.Time {
	return p.Started.Add(time.Minute * time.Duration(p.DurationMins))
}

func Start(fileName string) (pomodoro, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		ioutil.WriteFile(fileName, nil, 0644)
	}

	pomodoros, err := Load(fileName)
	if err != nil {
		return pomodoro{}, nil
	}

	newPomodoro := pomodoro{
		Started:      time.Now().UTC(),
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
