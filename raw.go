package json

import (
	"bytes"
	"errors"
	"unsafe"
)

type RawMessage []byte

func (m RawMessage) MarshalJSON() ([]byte, error) {
	if len(m) == 0 {
		return NullBytes, nil
	}
	if !Validate(m) {
		return nil, errors.New("json.RawMessage: MarshalJSON on invalid message")
	}
	return m, nil
}

func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	if data == nil || len(data) == 0 {
		return nil
	}
	if !Validate(data) {
		return errors.New("json.RawMessage: UnmarshalJSON on invalid message")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

func (m RawMessage) MarshalBinary() (data []byte, err error) {
	data, err = m.MarshalJSON()
	return
}

func (m *RawMessage) UnmarshalBinary(data []byte) (err error) {
	err = m.UnmarshalJSON(data)
	return
}

func (m RawMessage) TryMapToObject() bool {
	if m == nil {
		return false
	}
	if m[0] != '{' || m[len(m)-1] != '}' {
		return false
	}
	return true
}

func (m RawMessage) MapToObject() (r *Object, err error) {
	if !m.TryMapToObject() {
		err = errors.New("json.RawMessage: MapToObject on invalid message")
		return
	}
	r = &Object{
		raw: m,
	}
	return
}

func (m RawMessage) TryMapToArray() bool {
	if m == nil {
		return false
	}
	if m[0] != '[' || m[len(m)-1] != ']' {
		return false
	}
	return true
}

func (m RawMessage) MapToArray() (r *Array, err error) {
	if !m.TryMapToArray() {
		err = errors.New("json.RawMessage: MapToArray on invalid message")
		return
	}
	r = &Array{
		raw: m,
	}
	return
}

// Valid
// objects.Valid
func (m RawMessage) Valid() (ok bool) {
	ok = len(m) > 0 && !bytes.Equal(m, NullBytes)
	return
}

// Value
// objects.Value
func (m RawMessage) Value() (v any) {
	v = m
	return
}

// Marshal
// objects.Marshal
func (m RawMessage) Marshal() (p []byte, err error) {
	p = m
	return
}

// Unmarshal
// objects.Unmarshal
func (m RawMessage) Unmarshal(dst interface{}) (err error) {
	if !m.Valid() {
		return
	}
	if bytes.Equal(m, EmptyObjectBytes) || bytes.Equal(m, EmptyArrayBytes) {
		return
	}
	err = Unmarshal(m, dst)
	if err != nil {
		err = errors.New("json.RawMessage: transform failed")
		return
	}
	return
}

// Scan
// sql.Scanner
func (m *RawMessage) Scan(src interface{}) (err error) {
	switch s := src.(type) {
	case string:
		p := unsafe.Slice(unsafe.StringData(s), len(s))
		*m = bytes.Clone(p)
		break
	case []byte:
		*m = bytes.Clone(s)
		break
	case nil:
		break
	default:
		err = errors.New("json.RawMessage: scan failed, unsupported src type")
		break
	}
	return
}
