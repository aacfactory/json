package json_test

import (
	"fmt"
	"github.com/aacfactory/json"
	"testing"
	"time"
)

func TestRawMessage_UnmarshalJSON(t *testing.T) {
	d := time.Now()

	p, _ := json.Marshal(d)

	raw := json.RawMessage{}
	decodeErr := json.Unmarshal(p, &raw)
	t.Log(decodeErr, string(raw))
}

func TestRawMessage_MarshalJSON(t *testing.T) {

	p, mErr := json.Marshal(10)

	fmt.Println(p, mErr)

	fmt.Println(json.Marshal(p))

	fmt.Println(json.Unmarshal([]byte("null"), &json.RawMessage{}))

}
