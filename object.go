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
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var (
	EmptyObjectBytes = []byte{'{', '}'}
)

func IsEmptyObject(p []byte) (ok bool) {
	ok = len(p) == 0 || bytes.Equal(p, EmptyObjectBytes)
	return
}

func NewObjectFromInterface(v interface{}) *Object {
	if v == nil {
		panic(fmt.Errorf("new json object from interface failed, value is nil"))
	}
	b, encodeErr := Marshal(v)
	if encodeErr != nil {
		panic(fmt.Errorf("new json object from interface failed, encode value to json failed, %v", encodeErr))
	}
	return NewObjectFromBytes(b)
}

func NewObjectFromBytes(b []byte) *Object {
	if b[0] != '{' || b[len(b)-1] != '}' {
		panic(fmt.Errorf("new json object from bytes failed, %s is not json object bytes", string(b)))
	}
	return &Object{
		raw: b,
	}
}

func NewObject() *Object {
	return &Object{
		raw: []byte{'{', '}'},
	}
}

type Object struct {
	raw []byte
}

func (object *Object) Empty() (ok bool) {
	if object.raw == nil || len(object.raw) == 0 || bytes.Equal(object.raw, EmptyObjectBytes) {
		ok = true
		return
	}
	ok = !gjson.ParseBytes(object.raw).Exists()
	return
}

func (object *Object) Raw() (raw []byte) {
	raw = object.raw
	return
}

func (object *Object) WriteTo(out *Object) (err error) {
	if out == nil {
		err = fmt.Errorf("fns json.Object WriteTo: can not write to nil point")
		return
	}
	err = out.UnmarshalJSON(object.raw)
	return
}

func (object *Object) Merge(src *Object) (err error) {
	if src == nil {
		err = fmt.Errorf("fns json.Object Merge: can not write to nil point")
		return
	}
	dst := object.raw
	srcResult := gjson.ParseBytes(src.Raw())
	srcResult.ForEach(func(key gjson.Result, value gjson.Result) bool {
		dst = merge(dst, key.String(), value)
		return true
	})
	object.raw = dst
	return
}

func merge(dst []byte, srcKey string, srcValue gjson.Result) (result []byte) {
	switch srcValue.Type {
	case gjson.String, gjson.Number, gjson.True, gjson.False:
		affected, setErr := sjson.SetRawBytes(dst, srcKey, []byte(srcValue.Raw))
		if setErr != nil {
			result = dst
			return
		}
		result = affected
	case gjson.JSON:
		if srcValue.IsArray() {
			affected, setErr := sjson.SetRawBytes(dst, srcKey, []byte(srcValue.Raw))
			if setErr != nil {
				result = dst
				return
			}
			result = affected
			return
		}
		if srcValue.IsObject() {
			dstSub := gjson.GetBytes(dst, srcKey)
			if !dstSub.Exists() {
				affected, setErr := sjson.SetRawBytes(dst, srcKey, []byte(srcValue.Raw))
				if setErr != nil {
					result = dst
					return
				}
				result = affected
				return
			}

			dstSubRas := []byte(dstSub.Raw)
			srcValue.ForEach(func(key, value gjson.Result) bool {
				dstSubRas = merge(dstSubRas, key.Str, value)
				return true
			})

			affected, setErr := sjson.SetRawBytes(dst, srcKey, dstSubRas)
			if setErr != nil {
				result = dst
				return
			}
			result = affected

		}
	default:
		result = dst
	}

	return
}

func (object *Object) Contains(path string) (has bool) {
	has = gjson.GetBytes(object.raw, path).Exists()
	return
}

func (object *Object) Get(path string, v interface{}) (err error) {
	if path == "" {
		err = fmt.Errorf("json object get failed, path is empty")
		return
	}
	if v == nil {
		err = fmt.Errorf("json object get %s failed, value is nil", path)
		return
	}
	r := gjson.GetBytes(object.raw, path)
	if !r.Exists() {
		err = fmt.Errorf("json object get %s failed, not exists", path)
		return
	}
	decodeErr := Unmarshal([]byte(r.Raw), v)
	if decodeErr != nil {
		err = fmt.Errorf("json object get %s failed, decode failed", path)
		return
	}
	return
}

func (object *Object) Put(path string, v interface{}) (err error) {
	if path == "" {
		err = fmt.Errorf("json object set failed, path is empty")
		return
	}
	if v == nil {
		err = fmt.Errorf("json object set %s failed, value is nil", path)
		return
	}

	affected, setErr := sjson.SetBytes(object.raw, path, v)

	if setErr != nil {
		err = fmt.Errorf("json object set %s failed", path)
		return
	}
	object.raw = affected
	return
}

func (object *Object) PutRaw(path string, raw []byte) (err error) {
	if path == "" {
		err = fmt.Errorf("json object set raw failed, path is empty")
		return
	}
	if raw == nil || len(raw) == 0 {
		err = fmt.Errorf("json object set raw %s failed, value is nil", path)
		return
	}
	affected, setErr := sjson.SetRawBytes(object.raw, path, raw)
	if setErr != nil {
		err = fmt.Errorf("json object set %s failed", path)
		return
	}
	object.raw = affected
	return
}

func (object *Object) Remove(path string) (err error) {
	if path == "" {
		err = fmt.Errorf("json object remove failed, path is empty")
		return
	}

	affected, remErr := sjson.DeleteBytes(object.raw, path)
	if remErr != nil {
		err = fmt.Errorf("json object remove %s failed", path)
		return
	}
	object.raw = affected
	return
}

func (object *Object) MapTo(v interface{}) (err error) {
	err = Unmarshal(object.raw, v)
	return
}

func (object *Object) MarshalJSON() (p []byte, err error) {
	p = object.raw
	return
}

func (object *Object) UnmarshalJSON(p []byte) (err error) {
	if p == nil || len(p) == 0 {
		return
	}
	if p[0] != '{' || p[len(p)-1] != '}' {
		err = fmt.Errorf("unmarshal json object from bytes failed, %s is not json object bytes", string(p))
		return
	}
	object.raw = p
	return
}
