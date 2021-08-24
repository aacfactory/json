package json_test

import (
	"github.com/aacfactory/json"
	"testing"
	"time"
)

func TestRawMessage_UnmarshalJSON(t *testing.T) {
	d := time.Now()

	p, _ := json.Marshal(d)

	raw := json.RawMessage{}
	decodeErr :=json.Unmarshal(p, &raw)
	t.Log(decodeErr, string(raw))
}
