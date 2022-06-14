package selectors

import (
	"fmt"
	"time"
)

const (
	TimeLayout = time.RFC3339
)

// TimestampRange allows selecting events in a time range.
type TimestampRange struct {
	StartTimestamp Timestamp `query:"start_timestamp"`
	EndTimestamp   Timestamp `query:"end_timestamp"`
}

// Time represents a thin wrapper around `time.Time`. With a custom
// type defined we can easily enforce format validation by echo on
// binding query parameters.
type Timestamp time.Time

func (t *Timestamp) UnmarshalParam(param string) error {

	ts, err := time.Parse(TimeLayout, param)
	if err != nil {
		return fmt.Errorf("could not parse time: %w", err)
	}

	*t = Timestamp(ts)
	return nil
}
