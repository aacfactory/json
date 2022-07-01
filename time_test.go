package json_test

import (
	"fmt"
	"github.com/aacfactory/json"
	"testing"
	"time"
)

func TestTime_MarshalJSON(t *testing.T) {
	now := time.Now()
	p, err := json.Marshal(now)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(p))
	v := time.Time{}
	err = json.Unmarshal(p, &v)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(v)

	zero := time.Time{}
	p, err = json.Marshal(zero)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(p))
	v = time.Time{}
	err = json.Unmarshal(p, &v)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(v)

}
