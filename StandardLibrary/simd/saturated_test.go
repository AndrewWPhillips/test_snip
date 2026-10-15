package __test

import (
	"simd/archsimd"
	"testing"
)

// TestAddSaturated shows that after adding the two values they get clipped to the range -128 to 127 (inclusive)
func TestAddSaturated(t *testing.T) {
	x, y := RandVector8x32(0), RandVector8x32(0)
	result := x.AddSaturated(y)
	println(x.String())
	println(y.String())
	println(result.String())
}

func TestSaturateToInt8(t *testing.T) {
	// Unfortunately only AVX512 supports SaturateToInt8/16/32
	if !archsimd.X86.AVX512() {
		println("not AVX512")
		return
	}
	println(archsimd.LoadInt64x4Slice([]int64{42, 127, 300, -300}).SaturateToInt8().String())
}

func TestSaturateToInt16Concat(t *testing.T) {
	x := archsimd.LoadInt32x4Slice([]int32{1, 32767, 32768, 1e9})
	y := archsimd.LoadInt32x4Slice([]int32{-1, -32767, -32768, -32769})
	println(x.SaturateToInt16Concat(y).String())
}
