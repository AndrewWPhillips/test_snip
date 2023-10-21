package __

import (
	"testing"

	"golang.org/x/exp/constraints"
)

// tail adds the values in a slice using tail recursion
func tail(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return s[0] + tail(s[1:])
}

func BenchmarkTailRecursion(b *testing.B) { // ~100
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = tail([]int{1, 2, 3, 99})
	}
	println(sum)
}

// gtail is a generic version of tail that works with a slice of any integer type
func gtail[T constraints.Integer](s []T) T {
	if len(s) == 0 {
		return T(0)
	}
	return s[0] + gtail(s[1:])
}

func BenchmarkGenericTailRecursion(b *testing.B) { // ~250
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = gtail([]int{1, 2, 3, 99})
	}
	println(sum)
}

// iter adds the values in a slice using a loop
//
//go:noinline
func iter(s []int) int {
	r := 0
	for _, v := range s {
		r += v
	}
	return r
}

func BenchmarkIteration(b *testing.B) { // ~55 (or ~40 with inlining)
	sum := 0
	for i := 0; i < b.N; i++ {
		sum = iter([]int{1, 2, 3, 99})
	}
	println(sum)
}
