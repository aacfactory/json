package json

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

func TimeNow() Time {
	return Time(time.Now())
}

func NewTime(t time.Time) Time {
	return Time(t)
}

type Time time.Time

func (t Time) ToTime() time.Time {
	return time.Time(t)
}

func (t Time) String() string {
	return time.Time(t).Format(time.RFC3339)
}

func (t Time) MarshalJSON() (p []byte, err error) {
	if t.ToTime().IsZero() {
		p = []byte("\"\"")
		return
	}
	p = []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339)))
	return
}

func (t *Time) UnmarshalJSON(p []byte) error {
	if t == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	if p == nil {
		return nil
	}
	p = bytes.TrimSpace(p)
	if len(p) == 0 {
		return nil
	}
	if p[0] == '"' {
		if len(p) < 2 {
			return errors.New("json.RawMessage: UnmarshalJSON on invalid content")
		}
		p = p[1 : len(p)-1]
	}
	if len(p) == 0 {
		return nil
	}
	dt, parseErr := time.Parse(time.RFC3339, string(p))
	if parseErr != nil {
		return fmt.Errorf("parse time failed for invalid layout, raw value is %s", string(p))
	}
	*t = Time(dt)
	return nil
}
