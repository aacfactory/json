package json_test

import (
	"github.com/aacfactory/json"
	"math/big"
	"testing"
)

func TestBigRat(t *testing.T) {
	v := big.NewRat(1, 1)
	p, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(p))
	err = json.Unmarshal(p, v)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}

func TestBigInt(t *testing.T) {
	v := big.NewInt(1)
	p, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(p))
	err = json.Unmarshal(p, v)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}

func TestBigFloat(t *testing.T) {
	v := big.NewFloat(1.222)
	p, err := json.Marshal(v)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(p))
	err = json.Unmarshal(p, v)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}
