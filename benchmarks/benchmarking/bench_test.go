package __

// various tests of how optimization affects benchmarks
// Note: in GoLand I added "-test.count=10" to Program Args in the build config

import (
	"fmt"
	"log"
	"testing"
)

var TestGlobal any

// 7.0ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkNothing(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

// BenchmarkNothing1p24 uses new benchmarking of Go 1.24
func BenchmarkNothing1p24(b *testing.B) {
	for b.Loop() {
	}
}

// 7.0ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := 1
		_ = j
	}
}

// 7.05ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkUnderscore(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = 1
	}
	_ = j
}

// 7.2ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkGlobal(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = 1
	}
	TestGlobal = j
}

// 7.2ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkPrint(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = 1
	}
	println(j)
}

// 7.2ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkFmtPrint(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = 1
	}
	fmt.Println(j)
}

// 7.2ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkLogPrint(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = 1
	}
	log.Println(j)
}

func f(v int) int {
	return v / 2
}

//go:noinline
func noInline(v int) int {
	return v / 2
}

// 7.0ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := f(i)
		_ = j
	}
}

// 7.4ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := noInline(i)
		_ = j
	}
}

// 7.2ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkFuncKeep(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = f(i)
	}
	TestGlobal = j
}

// 7.6ns/op - Go 1.19 Windows/amd64 AMD Ryzen 5 2600 Six-Core Processor

func BenchmarkFuncNoInlineKeep(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = noInline(i)
	}
	TestGlobal = j
}
