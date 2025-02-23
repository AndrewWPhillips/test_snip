//go:build go1.18

package __

import (
	"fmt"
	"testing"
)

func BenchmarkAllocSmallInt(b *testing.B) {
	b.ReportAllocs()
	var x any
	for i := 0; i < b.N; i++ {
		a := 255
		escapeInt(&a)
		x = a
	}
	println(x)
}

func BenchmarkAllocLargeInt(b *testing.B) {
	b.ReportAllocs()
	var x any
	for i := 0; i < b.N; i++ {
		a := 256
		escapeInt(&a)
		x = a
	}
	fmt.Printf("%T %v\n", x, x)
}

var savedForBenchmarkAllocInt *int

func escapeInt(p *int) {
	savedForBenchmarkAllocInt = p
}
