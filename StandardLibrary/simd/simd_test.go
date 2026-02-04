package __test

import (
	"simd/archsimd"
	"testing"
)

func Add(a, b []float32) []float32 {
	if len(a) != len(b) {
		panic("slices of different length")
	}
	res := make([]float32, len(a))
	for i := range a {
		res[i] = a[i] + b[i]
	}
	return res
}

func SimdAdd(a, b []float32) []float32 {
	if len(a) != len(b) {
		panic("slices of different length")
	}

	// If AVX-512 isn't supported, fall back to scalar addition,
	// since the Float32x16.Add method needs the AVX-512 instruction set.
	if !archsimd.X86.AVX512() {
		return Add(a, b)
	}

	res := make([]float32, len(a))
	n := len(a)
	i := 0

	// 1. SIMD loop: Process 16 elements at a time.
	for i <= n-16 {
		// Load 16 elements from a and b vectors.
		va := archsimd.LoadFloat32x16Slice(a[i : i+16])
		vb := archsimd.LoadFloat32x16Slice(b[i : i+16])

		// Add all 16 elements in a single instruction
		// and store the results in the result vector.
		vSum := va.Add(vb) // translates to VADDPS asm instruction
		vSum.StoreSlice(res[i : i+16])

		i += 16
	}

	// 2. Scalar tail: Process any remaining elements (0-15).
	for ; i < n; i++ {
		res[i] = a[i] + b[i]
	}

	return res
}

var (
	a1 = []float32{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17,
	}
	a2 = []float32{
		17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2,
		1, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}
)

func BenchmarkSimdAdd(b *testing.B) {
	for b.Loop() {
		SimdAdd(a1, a2)
	}
}

func BenchmarkNormalAdd(b *testing.B) {
	for b.Loop() {
		Add(a1, a2)
	}
}
