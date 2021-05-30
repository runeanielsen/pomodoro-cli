package pomodoro_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/runeanielsen/pomodoro-cli/internal/pomodoro"
)

func TestStart(t *testing.T) {
	testCases := []struct {
		name           string
		expectErrMsg   string
		expectDuration int8
		expectedLen    int
	}{
		{"Start", "", 25, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tf, err := ioutil.TempFile("", "")

			if err != nil {
				t.Fatalf("Error creating temp file: %s", err)
			}
			defer os.Remove(tf.Name())

			p, err := pomodoro.Start(tf.Name())

			if tc.expectErrMsg != "" {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}

				if err.Error() != tc.expectErrMsg {
					t.Fatalf("Expected error message: '%s' got '%s' instead.", tc.expectErrMsg, err)
				}

				return
			}

			if tc.expectDuration != p.DurationMins {
				t.Fatalf("Expected DurationMins '%d' got '%d' instead", tc.expectDuration, p.DurationMins)
			}

			if p.Started.IsZero() {
				t.Fatalf("Started time should not be zero time")
			}

			pomodoros, err := pomodoro.Load(tf.Name())
			if err != nil {
				t.Fatal(err)
			}

			if tc.expectedLen != len(pomodoros) {
				t.Fatalf("Expected lengh to be %d instead got %d", tc.expectedLen, len(pomodoros))
			}
		})
	}
}
