package __test

import (
	"math/rand/v2"
	"simd/archsimd"
	"slices"
	"testing"
	"unsafe"
)

// TestSimdAbs just tests the AVX (SIMD) instruction that generates absolute values of 32 bytes
func TestSimdAbs(t *testing.T) {
	a := slices.Repeat([]int8{0, 1, -1, 9, -9, 127, -127, -128}, 8)
	aVec32 := archsimd.LoadInt8x32Slice(a)
	clear(a) // verify that changing the array doesn't change underlying values of aVec32

	result := aVec32.Abs()
	println(aVec32.String())
	println(result.String())
}

// RandInt8Vector generates a slice of 64 random signed bytes
// If the parameter (halfRange) is zero then we use the full range (-127 to 127 inclusive) but without -128
func RandInt8Vector(halfRange int) []int8 {
	if halfRange <= 0 || halfRange > 127 {
		halfRange = 127
	}
	topRange := 2*halfRange + 1 // 255 for halfRange 127

	const MAX_VEC = 64
	r := make([]int8, 0, MAX_VEC)
	for range MAX_VEC {
		r = append(r, int8(rand.IntN(topRange)-halfRange))
	}
	return r
}

// RandVector8x16 creates an int8 vector fill with random byte values
// The parameter halfRange determines the +/- range of values, but if
// zero, negative or more than 127 then the range -127 to 127 is used.
func RandVector8x16(halfRange int) archsimd.Int8x16 {
	return archsimd.LoadInt8x16Slice(RandInt8Vector(halfRange))
}

func RandVector8x32(halfRange int) archsimd.Int8x32 {
	return archsimd.LoadInt8x32Slice(RandInt8Vector(halfRange))
}

func RandVector8x64(halfRange int) archsimd.Int8x64 {
	return archsimd.LoadInt8x64Slice(RandInt8Vector(halfRange))
}

func TestSimdLoad(t *testing.T) {
	if !archsimd.X86.AVX() {
		println("AVX not supported")
		return
	}

	aVec16 := RandVector8x16(0)
	println(aVec16.String())
	println(aVec16.AsInt64x2().String())
	aVec32 := RandVector8x32(0)
	println(aVec32.String())
	if archsimd.X86.AVX512() {
		aVec64 := RandVector8x64(0)
		println(aVec64.String())
	}
}

func TestSimdCast(t *testing.T) {
	var p *archsimd.Int8x16
	v := struct{ vals [16]int8 }{
		[16]int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	p = (*archsimd.Int8x16)(unsafe.Pointer(&v))

	println(p.String())
}
