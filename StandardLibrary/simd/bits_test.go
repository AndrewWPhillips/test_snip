//go:build go1.26 && amd64 && goexperiment.simd

package __

import (
	"simd/archsimd"
	"testing"
)

func TestGetBitsFromMask(t *testing.T) {
	if !archsimd.X86.AVX2() {
		t.Fatal("AVX2 is required")
	}
	var v1 archsimd.Int16x16
	var b archsimd.Mask16x16
	b = v1.Equal(v1) // AVX2 - sets all bits on
	// println(b.ToBits()) // AVX512 - panics on my machine

	b.ToInt16x16().StoreSlice(make([]int16, 16)) // AVX2 - OK
}

func TestMultiple(t *testing.T) {
	if !archsimd.X86.AVX2() {
		t.Fatal("AVX2 is required")
	}
	var v1 archsimd.Int16x16        // VPXOR     Y0, Y0, Y0
	v2 := v1.Equal(v1).ToInt16x16() // VPCMPEQW  Y0, Y0, Y1
	v1 = v1.And(v2)                 // VPAND     Y1, Y0, Y0

	var r [16]int16
	v1.Store(&r)
	println(r[0])
}

func TestNot(t *testing.T) {
	// NOTE that VPXOR a value with itself clears all bits
	var v1 archsimd.Int8x32  // VPXOR Y0, Y0, Y0
	var v2 archsimd.Int16x16 // VPXOR Y1,Y1,Y1

	// Bitwise complement (NOT) is done by XOR with all bits on, where
	// turning all bits on is done by comparing a value to itself
	// NOTE that comparing a value with itself (using VPCMPEQB, VPCMPEQW etc) sets all bits
	v1 = v1.Not() // VPCMPEQB Y0, Y0, Y2; VPXOR Y0, Y2, Y0
	// Note that the VPCMPEQW instruction here is redundant as Y2 already has all bits on
	v2 = v2.Not() // VPCMPEQW Y1, Y1, Y2; VPXOR Y1, Y2, Y1

	var r [32]int8
	v1.Store(&r)
	v2.AsInt8x32().Store(&r)
	println(r[0])
}

func TestGalois(t *testing.T) {
	//var v archsimd.Uint8x64
	//v.GaloisFieldAffineTransformInverse()
}
