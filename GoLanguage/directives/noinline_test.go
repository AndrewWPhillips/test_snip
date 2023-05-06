package __

import (
	"log"
	"testing"
)

func inc(pi *int) {
	*pi++
}

// BenchmarkIncInline test how a benchmark performs on an inlined file (about 1.5 ns/op on work computer)
func BenchmarkIncInline(b *testing.B) {
	j := 0
	for i := 0; i < b.N; i++ {
		inc(&j)
	}
	log.Println(j)
}

//go:noinline
func incNoinline(pi *int) {
	*pi++
}

// BenchmarkIncNoInline does the same as BenchmarkIncInline but with a non-inlined func (1.7 ns/op)
func BenchmarkIncNoInline(b *testing.B) {
	j := 0
	for i := 0; i < b.N; i++ {
		incNoinline(&j)
	}
	log.Println(j)
}

// BenchmarkIncInline just runs a benchmark on nothing to ensure it's different to inlined func test (0.3 ns/op)
func BenchmarkControl(b *testing.B) {
	j := 0
	for i := 0; i < b.N; i++ {
	}
	log.Println(j)
}
