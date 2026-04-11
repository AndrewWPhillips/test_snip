//go:build go1.26 && goexperiment.simd && amd64

package __

import (
	"fmt"
	"math/rand/v2"
	"simd/archsimd"
	"testing"
)

// BenchmarkSimdNot tests flipping all the bits in a large array using the
// archsimd.Int8x32.Not() method which underneath uses archsimd.Int8x32.Xor()
func BenchmarkSimdNot(b *testing.B) { // ~65 ns/op
	var in, out [1024]int8 // value to invert (bitwise NOT)
	for i := range in {
		in[i] = int8(rand.IntN(255) - 128)
	}

	for b.Loop() {
		var v archsimd.Int8x32
		for i := 0; i < 1024; i += 32 {
			v = archsimd.LoadInt8x32Slice(in[i : i+32])
			v.Not().StoreSlice(out[i : i+32])
		}
	}
	fmt.Println(in[:10])
	fmt.Println(out[:10])
}

// BenchmarkSimdXor should be faster than BenchmarkSimdNot because it only creates
// the 'bitsOn' vector once, but for some reason it is 50% slower
func BenchmarkSimdXor(b *testing.B) {
	var in, out [1024]int8 // value to invert (bitwise NOT)
	for i := range in {
		in[i] = int8(rand.IntN(255) - 128)
	}

	for b.Loop() {
		var v archsimd.Int8x32
		bitsOn := v.Equal(v).ToInt8x32() // vector with all bits turned on (2ns)
		for i := 0; i < 1024; i += 32 {
			v = archsimd.LoadInt8x32Slice(in[i : i+32])
			v.Xor(bitsOn).StoreSlice(out[i : i+32])
		}
	}
	fmt.Println(in[:10])
	fmt.Println(out[:10])
}

func BenchmarkFabricateBitsSetVector(b *testing.B) { // 2ns/op
	for b.Loop() {
		var v archsimd.Int8x32
		_ = v.Equal(v).ToInt8x32()
	}
}
