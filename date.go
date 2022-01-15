package json

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

func DateNow() Date {
	return NewDateFromTime(time.Now())
}

func NewDate(year int, month time.Month, day int) Date {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return Date(t)
}

func NewDateFromTime(t time.Time) Date {
	x := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return Date(x)
}

type Date time.Time

func (d Date) ToTime() time.Time {
	return time.Time(d)
}

func (d Date) String() string {
	return d.ToTime().Format("2006-01-02")
}

func (d Date) MarshalJSON() (p []byte, err error) {
	if d.ToTime().IsZero() {
		p = []byte("\"\"")
		return
	}
	p = []byte(fmt.Sprintf("\"%s\"", d.String()))
	return
}

func (d *Date) UnmarshalJSON(p []byte) error {
	if d == nil {
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
	date, parseErr := time.Parse("2006-01-02", string(p))
	if parseErr != nil {
		return fmt.Errorf("parse date failed for not int, raw value is %s", string(p))
	}
	*d = Date(date)
	return nil
}
