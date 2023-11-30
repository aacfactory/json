/*
 * Copyright 2021 Wang Min Xiang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package json

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

var (
	_json    jsoniter.API
	_shorted jsoniter.API
)

func init() {
	jsoniter.RegisterTypeEncoderFunc("time.Time", datetimeTypeEncoderFunc, datetimeIsEmpty)
	jsoniter.RegisterTypeDecoderFunc("time.Time", datetimeTypeDecoderFunc)
	jsoniter.RegisterTypeEncoderFunc("json.Date", dateTypeEncoderFunc, dateIsEmpty)
	jsoniter.RegisterTypeDecoderFunc("json.Date", dateTypeDecoderFunc)
	jsoniter.RegisterTypeEncoderFunc("json.Time", timeTypeEncoderFunc, timeIsEmpty)
	jsoniter.RegisterTypeDecoderFunc("json.Time", timeTypeDecoderFunc)
	jsoniter.RegisterTypeEncoderFunc("complex64", complexTypeEncoderFunc, complexIsEmpty)
	jsoniter.RegisterTypeDecoderFunc("complex64", complexTypeDecoderFunc)
	jsoniter.RegisterTypeEncoderFunc("complex128", complexTypeEncoderFunc, complexIsEmpty)
	jsoniter.RegisterTypeDecoderFunc("complex128", complexTypeDecoderFunc)
	_json = jsoniter.Config{
		MarshalFloatWith6Digits:       true,
		EscapeHTML:                    false,
		ObjectFieldMustBeSimpleString: true,
	}.Froze()
	_shorted = jsoniter.Config{
		SortMapKeys: true,
		EscapeHTML:  true,
	}.Froze()
}

func Config(config jsoniter.Config) {
	_json = config.Froze()
}

func ConfigCompatibleWithStandardLibrary() {
	_json = jsoniter.ConfigCompatibleWithStandardLibrary
}

func Shorted() jsoniter.API {
	return _shorted
}

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

func Default() jsoniter.API {
	return _json
}

func Validate(data []byte) bool {
	return _json.Valid(data)
}

func ValidateString(data string) bool {
	return _json.Valid(unsafe.Slice(unsafe.StringData(data), len(data)))
}

func Marshal(v interface{}) (p []byte, err error) {
	p, err = Default().Marshal(v)
	return
}

func Unmarshal(data []byte, v interface{}) (err error) {
	err = Default().Unmarshal(data, v)
	return
}

func UnsafeMarshal(v interface{}) []byte {
	p, err := Default().Marshal(v)
	if err != nil {
		panic("json marshal object in unsafe mode failed")
		return nil
	}
	return p
}

func UnsafeUnmarshal(data []byte, v interface{}) {
	err := Default().Unmarshal(data, v)
	if err != nil {
		panic("json unmarshal object in unsafe mode failed")
		return
	}
	return
}
