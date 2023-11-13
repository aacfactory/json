package json

import "bytes"

var (
	NullBytes = []byte("null")
)

func IsNull(p []byte) (ok bool) {
	ok = len(p) == 0 || bytes.Equal(p, NullBytes)
	return
}
