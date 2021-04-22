package datetime

import (
	"time"
)

// GetDateStartFromString gets start of date from string
func GetDateStartFromString(d string) (*time.Time, error) {
	if d == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, d)
	if err != nil {
		return nil, err
	}

	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return &t, nil
}
