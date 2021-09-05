package json_test

import (
	"fmt"
	"github.com/aacfactory/json"
	"testing"
)

func TestTime_MarshalJSON(t *testing.T) {
	dt := json.TimeNow()

	p, err := json.Marshal(dt)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(p))
	td := json.Time{}
	err = json.Unmarshal(p, &td)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(td)
}
