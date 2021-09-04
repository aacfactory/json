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

	r := json.RawMessage{}

	fmt.Println(json.Marshal(r))



	fmt.Println(json.Unmarshal([]byte("null"), &json.RawMessage{}))

}
