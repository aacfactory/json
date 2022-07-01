package json

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"time"
	"unsafe"
)

func DateNow() Date {
	return NewDateFromTime(time.Now())
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func NewDateFromTime(t time.Time) Date {
	return NewDate(t.Year(), t.Month(), t.Day())
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) ToTime() time.Time {
	if d.Year < 1 {
		d.Year = 1
	}
	if d.Month < 1 {
		d.Month = 1
	}
	if d.Day < 1 {
		d.Day = 1
	}
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.Local)
}

func (d Date) IsZero() (ok bool) {
	ok = d.Year < 2 && d.Month < 2 && d.Day < 2
	return
}

func (d Date) String() string {
	return d.ToTime().Format("2006-01-02")
}

var (
	dateType = reflect.TypeOf(Date{})
)

func dateTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect.NewAt(dateType, ptr).Elem().Interface().(Date)
	if v.IsZero() {
		stream.WriteString("")
	} else {
		stream.WriteString(v.String())
	}
	return
}

func dateIsEmpty(ptr unsafe.Pointer) bool {
	return reflect.NewAt(dateType, ptr).Elem().Interface().(Date).IsZero()
}

func dateTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	str := iter.ReadString()
	if iter.Error != nil {
		return
	}
	if str == "" {
		return
	}
	v, parseErr := time.Parse("2006-01-02", str)
	if parseErr != nil {
		iter.ReportError("unmarshal json.Date", parseErr.Error())
		return
	}
	reflect.NewAt(dateType, ptr).Elem().Set(reflect.ValueOf(NewDateFromTime(v)))
	return
}
