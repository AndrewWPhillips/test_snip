package __

import (
	"testing"
)

//go:noinline
func returnPointer() *int {
	i := 42
	return &i // i escapes so needs to be on the heap
}

func BenchmarkReturnPointer(b *testing.B) { // 14.32 ns/op  1 allocs/op
	b.ReportAllocs()
	outer := 0
	for i := 0; i < b.N; i++ {
		outer = *(returnPointer())
	}
	TestGlobal = outer
}

var staticInt = 42

//go:noinline
func returnStatic() *int {
	return &staticInt
}

func BenchmarkReturnStatic(b *testing.B) { // 1.6 ns/op  0 allocs/op
	b.ReportAllocs()
	outer := 0
	for i := 0; i < b.N; i++ {
		outer = *(returnStatic())
	}
	TestGlobal = outer
}

//go:noinline
func returnInt() int {
	return 42
}

func BenchmarkReturnInt(b *testing.B) { // 1.6 ns/op  0 allocs/op
	b.ReportAllocs()
	outer := 0
	for i := 0; i < b.N; i++ {
		outer = returnInt()
	}
	TestGlobal = outer
}
