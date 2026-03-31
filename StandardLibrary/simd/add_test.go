package __test

import (
	"simd/archsimd"
	"testing"
)

func TestSimdAdd(t *testing.T) {
	x := archsimd.LoadInt8x32(&[32]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 127, -127, 127, 127})
	y := archsimd.LoadInt8x32(&[32]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 127, 127, -128, 1})

	result := x.Add(y)
	println(result.String())
}

// TestSimdAddSaturated tests saturated add (
func TestSimdAddSaturated(t *testing.T) {
	x := archsimd.LoadInt8x32(&[32]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 127, -127, 127, 127})
	y := archsimd.LoadInt8x32(&[32]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 127, 127, -128, 1})

	result := x.AddSaturated(y)
	println(result.String())
}
