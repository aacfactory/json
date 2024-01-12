package json

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"time"
	"unsafe"
)

var (
	durType = reflect.TypeOf(time.Duration(0))
)

func durationTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect.NewAt(durType, ptr).Elem().Interface().(time.Duration)
	if v < 1 {
		stream.WriteString("")
	} else {
		stream.WriteString(v.String())
	}
	return
}

func durationIsEmpty(ptr unsafe.Pointer) bool {
	return reflect.NewAt(durType, ptr).Elem().Interface().(time.Duration) == 0
}

func durationTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	str := iter.ReadString()
	if iter.Error != nil {
		return
	}
	if str == "" {
		return
	}
	v, parseErr := time.ParseDuration(str)
	if parseErr != nil {
		iter.ReportError("unmarshal time.Duration", parseErr.Error())
		return
	}
	reflect.NewAt(durType, ptr).Elem().Set(reflect.ValueOf(v))
	return
}
