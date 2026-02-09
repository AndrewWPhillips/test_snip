package __

import (
	"fmt"
	"math/rand/v2"
	"simd/archsimd"
	"testing"
)

func BenchmarkSimdLoadAndAbs(b *testing.B) { // 1.78 ns/op
	a := make([]int8, 32)
	for i := range a {
		a[i] = int8(rand.IntN(255) - 127) // -127 => +127 (no -128)
	}
	for b.Loop() {
		aVec32 := archsimd.LoadInt8x32Slice(a[:])
		aVec32.Abs()
	}
}

func BenchmarkSimdJustAbs(b *testing.B) { // 1.75 ns/op
	a := make([]int8, 32)
	for i := range a {
		a[i] = int8(rand.IntN(255) - 127) // -127 => +127 (no -128)
	}
	aVec32 := archsimd.LoadInt8x32Slice(a[:])
	var result archsimd.Int8x32
	for b.Loop() {
		result = aVec32.Abs()
	}
	println(result.String())
}

func BenchmarkNormalAbs(b *testing.B) { // 33 ns/op
	a := make([]int8, 32)
	var r [32]int8
	for i := range a {
		a[i] = int8(rand.IntN(255) - 127) // -127 => +127 (no -128)
	}
	for b.Loop() {
		for i := range a {
			if a[i] >= 0 {
				r[i] = a[i]
			} else {
				r[i] = -a[i]
			}
		}
	}
	fmt.Println(r)
}
