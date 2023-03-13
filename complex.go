package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strconv"
	"unsafe"
)

var (
	complexType = reflect.TypeOf(complex128(0))
)

func complexTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect.NewAt(complexType, ptr).Elem().Complex()
	stream.WriteString(fmt.Sprintf("%v", v))
	return
}

func complexIsEmpty(ptr unsafe.Pointer) bool {
	return reflect.NewAt(complexType, ptr).Elem().Complex() == 0
}

func complexTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	str := iter.ReadString()
	if iter.Error != nil {
		return
	}
	if str == "" {
		return
	}
	v, parseErr := strconv.ParseComplex(str, 128)
	if parseErr != nil {
		iter.ReportError("unmarshal complex", parseErr.Error())
		return
	}
	reflect.NewAt(complexType, ptr).Elem().SetComplex(v)
	return
}
