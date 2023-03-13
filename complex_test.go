package json_test

import (
	"fmt"
	"github.com/aacfactory/json"
	"testing"
)

func TestComplex(t *testing.T) {
	c64 := complex64(complex(1, 8))
	p64, e64 := json.Marshal(c64)
	fmt.Println(string(p64), e64)
	c64e := complex64(complex(0, 0))
	e64 = json.Unmarshal(p64, &c64e)
	fmt.Println(e64, c64 == c64e)
}
