package json_test

import (
	"fmt"
	"github.com/aacfactory/json"
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	d := json.NewDate(2021, time.March, 12)
	p, err := json.Marshal(d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(p))
	x := json.Date{}
	err = json.Unmarshal(p, &x)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(x.String())
}
