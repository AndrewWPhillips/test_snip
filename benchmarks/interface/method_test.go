package __

import (
	"testing"
)

type (
	A int
	I interface {
		f()
	}
)

//go:noinline
func (a A) f() {
}

func BenchmarkDirectCall(b *testing.B) { // 2.2 ns/op
	var a A
	for i := 0; i < b.N; i++ {
		a.f()
	}
}

func BenchmarkInterfaceCall(b *testing.B) { // 2.8 ns/op
	var ii I
	var a A
	ii = a
	for i := 0; i < b.N; i++ {
		ii.f()
	}
}

func BenchmarkDevirtualizationCall(b *testing.B) {
	var ii I
	var a A
	ii = a
	for i := 0; i < b.N; i++ {
		if aa, ok := ii.(A); ok {
			aa.f()
		}
	}
}
