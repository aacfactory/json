package json

import "errors"

type RawMessage []byte

func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
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
	if !Validate(data) {
		return errors.New("json.RawMessage: UnmarshalJSON on invalid message")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

func (m *RawMessage) TryMapToObject() bool {
	if m == nil {
		return false
	}
	b := []byte(*m)
	if b[0] != '{' || b[len(b)-1] != '}' {
		return false
	}
	return true
}

func (m *RawMessage) MapToObject() (r *Object, err error) {
	if !m.TryMapToObject() {
		err = errors.New("json.RawMessage: MapToObject on invalid message")
		return
	}
	r = &Object{
		raw: *m,
	}
	return
}

func (m *RawMessage) TryMapToArray() bool {
	if m == nil {
		return false
	}
	b := []byte(*m)
	if b[0] != '[' || b[len(b)-1] != ']' {
		return false
	}
	return true
}

func (m *RawMessage) MapToArray() (r *Array, err error) {
	if !m.TryMapToArray() {
		err = errors.New("json.RawMessage: MapToArray on invalid message")
		return
	}
	r = &Array{
		raw: *m,
	}
	return
}
