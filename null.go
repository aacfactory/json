package json

import "bytes"

var (
	Null = []byte("null")
)

func IsNull(p []byte) (ok bool) {
	ok = len(p) == 0 || bytes.Equal(p, Null)
	return
}
