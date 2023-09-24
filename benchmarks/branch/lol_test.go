package __

import (
	"testing"
)

var TestGlobal any

// from https://medium.com/@ludirehak/printing-lol-doubled-the-speed-of-my-go-code-e32e02fc3f92

func if_max(values []int) int {
	maxV := values[0]
	for _, v := range values[1:] {
		if v > maxV {
			maxV = v
			continue
		}
		print("lol")
	}
	return maxV
}

func BenchmarkLOL(b *testing.B) {
	v := 0
	s := []int{2, 5, 7, 8, 11, 12}
	for i := 0; i < b.N; i++ {
		v = if_max(s)
	}
	TestGlobal = v
}

func BenchmarkLOLopt(b *testing.B) {
	s := []int{2, 5, 7, 8, 11, 12}
	for i := 0; i < b.N; i++ {
		_ = if_max(s)
	}
}
