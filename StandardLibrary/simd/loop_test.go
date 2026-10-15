//go:build go1.26 && goexperiment.simd && amd64

package __

import (
	"fmt"
	"math/rand/v2"
	"simd/archsimd"
	"strconv"
	"testing"
	"unsafe"
)

func TestSimdLoop(t *testing.T) {
	const SIMDVecSize = 32
	in := MakeInt8Slice(50+rand.IntN(5000), SIMDVecSize)
	out := make([]int8, len(in), cap(in))

	for i := 0; i < len(in); i += SIMDVecSize {
		vec := archsimd.LoadInt8x32Slice(in[i : i+SIMDVecSize])
		a := vec.Abs()
		a.StoreSlice(out[i : i+SIMDVecSize])
	}
	fmt.Println(len(in), cap(in), in[:32])
	fmt.Println(len(out), cap(out), out[:32])
}

// BenchmarkSimdLoop tests running over a large slice doing a SIMD 256-bit vector at a time.
// The SIMD operation is archsimd.Int8x32.Abs() which gets the absolute value of a signed byte.
// On my computer this took about 800 nanoseconds/op
func BenchmarkSimdLoop(b *testing.B) {
	const SIMDVecSize = 32 // Int8x32 has 256 bits or 32 bytes
	in := MakeInt8Slice(10000, SIMDVecSize)
	out := make([]int8, len(in), cap(in))

	for b.Loop() {
		for i := 0; i < len(in); i += SIMDVecSize {
			vec := archsimd.LoadInt8x32Slice(in[i : i+SIMDVecSize])
			a := vec.Abs()
			a.StoreSlice(out[i : i+SIMDVecSize])
		}
	}
	fmt.Println(len(in), cap(in), in[:32])
	fmt.Println(len(out), cap(out), out[:32])
}

// BenchmarkNormalLoop does the same as BenchmarkSimdLoop but without using SIMD
// On my computer this took about 36 microseconds/op, ie. 45X slower
func BenchmarkNormalLoop(b *testing.B) {
	in := MakeInt8Slice(10000, 32)
	out := make([]int8, len(in), cap(in))

	for b.Loop() {
		for i := range in {
			if in[i] < 0 {
				out[i] = -in[i]
			} else {
				out[i] = in[i]
			}
		}
	}
	fmt.Println(len(in), cap(in), in[:32])
	fmt.Println(len(out), cap(out), out[:32])
}

// MakeInt8Slice creates a large vector of given size with random values
// The returned slice is aligned on a 'pad'sized word boundary.
// The returned slice is also padded (using its capacity) up to the next 'pad' boundary
// so that SIMD instructions can run off the end (size) of the returned slice.
// Parameters:
// - size = number of elements (length of returned slice)
// - pad = size of underlying SIMD vector (16, 32 or 64)
func MakeInt8Slice(size, pad int) []int8 {
	if !(pad == 16 || pad == 32 || pad == 64) {
		panic("MakeInt8Slice invalid pad" + strconv.Itoa(pad))
	}
	// Make a buffer with extra space for
	// - alignment at the start on SIMD vector
	// - padding up to SIMD vector size at the end
	buf := make([]int8, size+pad*2)

	// Work out which bit of the buffer is to be used
	rem := int(uintptr(unsafe.Pointer(&buf[0]))) % pad
	start := (pad - rem) % pad
	end := ((size + pad - 1) / pad) * pad

	for i := start; i < start+size; i++ {
		buf[i] = int8(rand.IntN(256) - 128)
	}
	return buf[start : start+size : start+end]
}

func TestMakeInt8Vector(t *testing.T) {
	testData := map[string]struct{ in, expCap int }{
		"zero":  {0, 0},
		"one":   {1, 32},
		"less":  {31, 32},
		"exact": {32, 32},
		"more":  {33, 2 * 32},
		"more2": {63, 2 * 32},
		"more3": {64, 2 * 32},
		"more4": {65, 3 * 32},
	}

	for name, data := range testData {
		t.Run(name, func(t *testing.T) {
			out := MakeInt8Slice(data.in, 32)
			if len(out) != data.in {
				t.Fatalf("%-10s: expected len %d but got %d\n", name, data.in, len(out))
			}
			if cap(out) != data.expCap {
				t.Fatalf("%-10s: expected cap %d but got %d\n", name, data.expCap, cap(out))
			}
		})
	}
}
