package json

import (
	"errors"
	"fmt"
	"time"
)

func TimeNow() Time {
	return Time(time.Now())
}

type Time time.Time

func (t Time) String() string {
	return time.Time(t).Format(time.RFC3339)
}

func (t Time) MarshalJSON() (p []byte, err error) {
	p = []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339)))
	return
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}

	dt, parseErr := time.Parse(time.RFC3339, string(data))
	if parseErr != nil {
		return fmt.Errorf("parse time failed for invalid layout, raw value is %s", string(data))
	}
	*t = Time(dt)
	return nil
}
