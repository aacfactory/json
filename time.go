package json

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"time"
	"unsafe"
)

var (
	timeType = reflect.TypeOf(Time{})
)

type Time struct {
	Hour    int
	Minutes int
	Second  int
}

func (t Time) IsZero() (ok bool) {
	ok = t.Hour == 0 && t.Minutes == 0 && t.Second == 0
	return
}

func (t Time) ToTime() time.Time {
	if t.Hour < 0 || t.Hour > 23 {
		t.Hour = 0
	}
	if t.Minutes < 0 || t.Minutes > 59 {
		t.Minutes = 0
	}
	if t.Second < 0 || t.Second > 59 {
		t.Second = 0
	}
	return time.Date(1, 1, 1, t.Hour, t.Minutes, t.Second, 0, time.Local)
}

func (t Time) String() string {
	return t.ToTime().Format("15:04:05")
}

func timeTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect.NewAt(timeType, ptr).Elem().Interface().(Time)
	if v.IsZero() {
		stream.WriteString("")
	} else {
		stream.WriteString(v.ToTime().Format("15:04:05"))
	}
	return
}

func timeIsEmpty(ptr unsafe.Pointer) bool {
	return reflect.NewAt(timeType, ptr).Elem().Interface().(Time).IsZero()
}

func timeTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	str := iter.ReadString()
	if iter.Error != nil {
		return
	}
	if str == "" {
		return
	}
	v, parseErr := time.Parse("15:04:05", str)
	if parseErr != nil {
		iter.ReportError("unmarshal json.Time", parseErr.Error())
		return
	}
	reflect.NewAt(timeType, ptr).Elem().Set(reflect.ValueOf(v))
	return
}

var (
	datetimeType = reflect.TypeOf(time.Time{})
)

func datetimeTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect.NewAt(datetimeType, ptr).Elem().Interface().(time.Time)
	if v.IsZero() {
		stream.WriteString("")
	} else {
		stream.WriteString(v.Format(time.RFC3339))
	}
	return
}

func datetimeIsEmpty(ptr unsafe.Pointer) bool {
	return reflect.NewAt(datetimeType, ptr).Elem().Interface().(time.Time).IsZero()
}

func datetimeTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	str := iter.ReadString()
	if iter.Error != nil {
		return
	}
	if str == "" {
		return
	}
	v, parseErr := time.Parse(time.RFC3339, str)
	if parseErr != nil {
		iter.ReportError("unmarshal time.Time", parseErr.Error())
		return
	}
	reflect.NewAt(datetimeType, ptr).Elem().Set(reflect.ValueOf(v))
	return
}
