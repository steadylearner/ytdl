package ytdl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func TestReverseString(t *testing.T) {
	str := []byte("example")
	reverseByteSlice(str)

	assert.Equal(t, "elpmaxe", string(str))
}

func TestInterfaceToString(t *testing.T) {

	values := map[interface{}]string{
		8:         "8",
		int64(16): "16",
		"hello":   "hello",
		0.01:      "0.01",
	}

	for k, v := range values {
		if interfaceToString(k) != v {
			t.Error("Value:", v, " != ", v)
		}
	}

}
