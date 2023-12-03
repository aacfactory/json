package json_test

import (
	stdjson "encoding/json"
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

	fmt.Println(p, string(p), mErr)
	n := 0
	fmt.Println(json.Unmarshal(p, &n))
	fmt.Println(json.Marshal(p))

	fmt.Println(json.Unmarshal([]byte("null"), &json.RawMessage{}))

}

func TestValidate(t *testing.T) {
	p, _ := json.Marshal(1)
	fmt.Println(string(p), json.Validate(p), stdjson.Valid(p))
}
