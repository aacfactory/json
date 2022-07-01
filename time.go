package json

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"time"
	"unsafe"
)

func TimeNow() Time {
	return Time(time.Now())
}

func NewTime(t time.Time) Time {
	return Time(t)
}

// Time
// Deprecated
type Time time.Time

func (t Time) ToTime() time.Time {
	return time.Time(t)
}

func (t Time) String() string {
	return time.Time(t).Format(time.RFC3339)
}

var (
	timeType = reflect.TypeOf(time.Time{})
)

func timeTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect.NewAt(timeType, ptr).Elem().Interface().(time.Time)
	if v.IsZero() {
		stream.WriteString("")
	} else {
		stream.WriteString(v.Format(time.RFC3339))
	}
	return
}

func timeIsEmpty(ptr unsafe.Pointer) bool {
	return reflect.NewAt(timeType, ptr).Elem().Interface().(time.Time).IsZero()
}

func timeTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
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
	reflect.NewAt(timeType, ptr).Elem().Set(reflect.ValueOf(v))
	return
}
