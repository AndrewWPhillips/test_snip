package __

import (
	"testing"
)

// Use "NoBoundsCheck" configuration to run the tests without bounds checking (makes them 0.2ns = 5% faster)

// var v [9]int
var v = make([]int, 9)

var A, B, C, D, E, F, G, H, I int

func BenchmarkBoundsCheckInOrder(b *testing.B) {
	v := make([]int, 9)
	for n := 0; n < b.N; n++ {
		A = v[0]
		B = v[1]
		C = v[2]
		D = v[3]
		E = v[4]
		F = v[5]
		G = v[6]
		H = v[7]
		I = v[8]
	}
}

func BenchmarkBoundsCheckOutOfOrder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		I = v[8]
		A = v[0]
		B = v[1]
		C = v[2]
		D = v[3]
		E = v[4]
		F = v[5]
		G = v[6]
		H = v[7]
	}
}
