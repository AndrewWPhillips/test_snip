package __

import (
	"slices"
	"testing"

	"golang.org/x/exp/constraints"
)

// TestGlobal is just here for assigning calculated values to.
// This is used in benchmarks to ensure that code is not optimised away, since if you assign a value to
// a variable visible outside the package the compiler doesn't know it's not used anywhere.
var TestGlobal any

func BenchmarkMin(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = min(99, -2)
	}
	TestGlobal = j
}

func gmin[T constraints.Ordered](s ...T) T {
	r := s[0]
	for _, v := range s[1:] {
		if v < r {
			r = v
		}
	}
	return r
}

func BenchmarkGenericMin(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = gmin(99, -2)
	}
	TestGlobal = j
}

func gmin2[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func BenchmarkGenericMin2(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = gmin2(99, -2)
	}
	TestGlobal = j
}

func BenchmarkGenericMin3(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = slices.Min([]int{99, -2})
	}
	TestGlobal = j
}
