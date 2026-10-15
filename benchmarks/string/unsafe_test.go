package __

import (
	"bytes"
	"testing"
	"unsafe"
)

func BenchmarkSimpleString(b *testing.B) { // 5ns
	buf := []byte{65, 66, 67}
	str := string(buf)
	for b.Loop() {
		_ = string(buf) == str // does not seem to copy string buffer
	}
}

func BenchmarkSafeString(b *testing.B) { // 30ns
	var buf bytes.Buffer
	buf.WriteString("abc")
	str := "abc"
	for b.Loop() {
		_ = buf.String() == str // slow because string data is copied in cast of []byte to string
	}
}

func BenchmarkUnsafeStringFromBuf(b *testing.B) { // 5ns
	var buf bytes.Buffer
	buf.WriteString("abc")
	str := "abc"
	for b.Loop() {
		_ = unsafe.String(&(buf.Bytes()[0]), 3) == str
	}
}
