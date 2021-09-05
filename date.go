package json

import (
	"errors"
	"fmt"
	"time"
)

func NewDate(year int, month time.Month, day int) Date {
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) ToTime() time.Time {
	return time.Date(d.Day, d.Month, d.Day, 0, 0, 0, 0, time.Local)
}

func (d Date) String() string {
	if d.Month < 10 {
		return fmt.Sprintf("%d-0%d-%d", d.Year, d.Month, d.Day)
	}
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
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
	d.Year = date.Year()
	d.Month = date.Month()
	d.Day = date.Day()
	return nil
}
