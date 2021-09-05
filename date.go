package json

import (
	"errors"
	"fmt"
	"time"
)

func NewDate(year int, month time.Month, day int) Date {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return Date(t)
}

type Date time.Time

func (d Date) ToTime() time.Time {
	return time.Time(d)
}

func (d Date) String() string {
	return d.ToTime().Format("2006-01-02")
}

func (d Date) MarshalJSON() (p []byte, err error) {
	p = []byte(fmt.Sprintf("\"%s\"", d.String()))
	return
}

func (d *Date) UnmarshalJSON(data []byte) error {
	if d == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}

	date, parseErr := time.Parse("2006-01-02", string(data))
	if parseErr != nil {
		return fmt.Errorf("parse date failed for not int, raw value is %s", string(data))
	}
	*d = Date(date)
	return nil
}
