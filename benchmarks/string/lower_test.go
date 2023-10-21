package __

import (
	"testing"
)

var Global interface{}

// lower returns the ASCII lowercase version of b
// This is from std lib (see strings/internal/ascii/print.go)
func lower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func BenchmarkStdStringInternalAscii(b *testing.B) { // ~7.6ns/op
	bb := byte('N')
	for i := 0; i < b.N; i++ {
		bb = lower(bb)
	}
	Global = bb
}

// lower(c) is a lower-case letter if and only if
// c is either that lower-case letter or the equivalent upper-case letter.
// Instead of writing c == 'x' || c == 'X' one can write lower(c) == 'x'.
// Note that lower of non-letters can produce other non-letters.
func strconv_lower(c byte) byte {
	return c | ('x' - 'X')
}

func BenchmarkStdStrconv(b *testing.B) { // ~7.6ns/op
	bb := byte('N')
	for i := 0; i < b.N; i++ {
		bb = strconv_lower(bb)
	}
	Global = bb
}

func myLower(b byte) byte {
	if b-'A' <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func BenchmarkMine(b *testing.B) { // ~7.6ns/op
	bb := byte('N')
	for i := 0; i < b.N; i++ {
		bb = myLower(bb)
	}
	Global = bb
}
