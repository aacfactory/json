package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"math/big"
	"reflect"
	"unsafe"
)

var (
	bigRatType   = reflect.TypeOf(big.NewRat(1, 1)).Elem()
	bigIntType   = reflect.TypeOf(big.NewInt(1)).Elem()
	bigFloatType = reflect.TypeOf(big.NewFloat(1)).Elem()
)

func bigRatTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect2.Type2(bigRatType).PackEFace(ptr).(*big.Rat)
	text, _ := v.MarshalText()
	_, _ = stream.Write(append(append([]byte{'"'}, text...), '"'))
	return
}

func bigRatIsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func bigRatTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.Error != nil {
		return
	}
	str := iter.ReadString()
	if str == "" {
		return
	}
	p := unsafe.Slice(unsafe.StringData(str), len(str))
	v := reflect2.Type2(bigRatType).PackEFace(ptr).(*big.Rat)
	parseErr := v.UnmarshalText(p)
	if parseErr != nil {
		iter.ReportError("unmarshal big.Rat", parseErr.Error())
		return
	}
	return
}

func bigIntTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect2.Type2(bigIntType).PackEFace(ptr).(*big.Int)
	text, _ := v.MarshalText()
	_, _ = stream.Write(append(append([]byte{'"'}, text...), '"'))
	return
}

func bigIntIsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func bigIntTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.Error != nil {
		return
	}
	str := iter.ReadString()
	if str == "" {
		return
	}
	p := unsafe.Slice(unsafe.StringData(str), len(str))
	v := reflect2.Type2(bigIntType).PackEFace(ptr).(*big.Int)
	parseErr := v.UnmarshalText(p)
	if parseErr != nil {
		iter.ReportError("unmarshal big.Rat", parseErr.Error())
		return
	}
	return
}

func bigFloatTypeEncoderFunc(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := reflect2.Type2(bigFloatType).PackEFace(ptr).(*big.Float)
	text, _ := v.MarshalText()
	_, _ = stream.Write(append(append([]byte{'"'}, text...), '"'))
	return
}

func bigFloatIsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func bigFloatTypeDecoderFunc(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.Error != nil {
		return
	}
	str := iter.ReadString()
	if str == "" {
		return
	}
	p := unsafe.Slice(unsafe.StringData(str), len(str))
	v := reflect2.Type2(bigFloatType).PackEFace(ptr).(*big.Float)
	parseErr := v.UnmarshalText(p)
	if parseErr != nil {
		iter.ReportError("unmarshal big.Rat", parseErr.Error())
		return
	}
	return
}
