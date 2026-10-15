package __

import (
	"strings"
	"testing"
	"unsafe"
)

// TestFromByteArray tests making a string from a byte array/slice
func TestFromByteArray(t *testing.T) {
	b := []byte{65, 66}

	// This gives 3 different addresses - ie the string data is copied
	println(b)                            // [2/2]0x2c456ab1f56
	println(unsafe.StringData(string(b))) // 0x2c456ab1f36
	println(unsafe.StringData(string(b))) // 0x2c456ab1f16

	// This is the efficient but unsafe way to create a string from a byte array
	// It's unsafe because changing b will confuse things
	println(unsafe.StringData(unsafe.String(&b[0], len(b)))) // 0x2c456ab1f56

	sb := strings.Builder{}
	sb.Write(b)
	println(unsafe.StringData(sb.String())) // heap 0xc00000a450
}

func TestHeapOrStack(t *testing.T) {
	b2 := make([]byte, 32) // stack
	println(unsafe.StringData(string(b2)))
	b2 = make([]byte, 33) // heap
	println(unsafe.StringData(string(b2)))
	b2 = make([]byte, 256)
	println(unsafe.StringData(string(b2)))
}

func TestBook(t *testing.T) {
}
