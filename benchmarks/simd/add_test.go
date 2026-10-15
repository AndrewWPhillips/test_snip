package __

import (
	"simd/archsimd"
	"testing"
)

func BenchmarkSimdAdd(b *testing.B) { // 1.75 ns/op
	x := archsimd.LoadInt8x32(&[32]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 127, -127, 127, 127})
	y := archsimd.LoadInt8x32(&[32]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 127, 127, -128, 1})
	var result archsimd.Int8x32
	for b.Loop() {
		result = x.Add(y)
	}
	println(result.String())
}

func BenchmarkSimdLoadAndAdd(b *testing.B) { // 1.71 ns/op
	x := [32]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 127, -127, 127, 127}
	y := [32]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 127, 127, -128, 1}
	var result archsimd.Int8x32
	for b.Loop() {
		result = archsimd.LoadInt8x32(&x).Add(archsimd.LoadInt8x32(&y))
	}
	println(result.String())
}

func BenchmarkNormalAdd(b *testing.B) { // 27 ns/op
	x := [32]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 127, -127, 127, 127}
	y := [32]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 127, 127, -128, 1}
	var result [32]int8
	for b.Loop() {
		for i := range 32 {
			result[i] = x[i] + y[i]
		}
	}
	println(result[0], result[1])
}
