package events

import (
	"fmt"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04"
)

// TimeSelector allows selecting events in a time range.
type TimeSelector struct {
	Start Time `query:"start"`
	End   Time `query:"end"`
}

// Time represents a thin wrapper around `time.Time`. With a custom
// type defined we can easily enforce format validation by echo on
// binding query parameters.
type Time time.Time

func (t *Time) UnmarshalParam(param string) error {

	ts, err := time.Parse(TimeLayout, param)
	if err != nil {
		return fmt.Errorf("could not parse time: %w", err)
	}

	*t = Time(ts)
	return nil
}
