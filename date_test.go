package json_test

import (
	"fmt"
	"github.com/aacfactory/json"
	"reflect"
	"testing"
)

func TestNewDate(t *testing.T) {
	fmt.Println(reflect.TypeOf(json.Date{}).String())
	now := json.DateNow()
	p, err := json.Marshal(now)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(p))
	v := json.Date{}
	fmt.Println(v)
	err = json.Unmarshal(p, &v)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
