package json_test

import (
	"github.com/aacfactory/json"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	dur := 1 * time.Hour
	p, encodeErr := json.Marshal(dur)
	if encodeErr != nil {
		t.Error(encodeErr)
		return
	}
	t.Log(string(p))
	dur = time.Duration(0)
	decodeErr := json.Unmarshal(p, &dur)
	if decodeErr != nil {
		t.Error(decodeErr)
		return
	}
	t.Log(dur)
}
