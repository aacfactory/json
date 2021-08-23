package json_test

import (
	"github.com/aacfactory/json"
	"testing"
	"time"
)

func TestNewObject(t *testing.T) {
	obj := json.NewObject()
	_ = obj.Put("hello", "world")
	_ = obj.Put("now", time.Now())
	_ = obj.Put("latency", 123*time.Millisecond)

	t.Log(obj.Contains("hello"))

	p, encodeErr := json.Marshal(obj)
	t.Log(string(p), encodeErr)

	o := json.NewObject()

	decodeErr := json.Unmarshal(p, o)
	t.Log(decodeErr)
	world := ""
	_ = o.Get("hello", &world)
	t.Log(world)
	now := time.Time{}
	_ = o.Get("now", &now)
	t.Log(now)
	latency := time.Duration(0)
	_ = o.Get("latency", &latency)
	t.Log(latency)

}
