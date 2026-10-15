//go:build go1.26 && goexperiment.simd && amd64

package __

import (
	"reflect"
	"simd/archsimd"
	"slices"
	"testing"
)

func TestCutOff(t *testing.T) {
	const CUTOFF = 5
	in := []int16{0, 1, 2, 3, 4, CUTOFF, 6, 7, 8, 9, 10, 11, 12, 13, 100, 1000}
	exp := []int16{0, 2, 4, 6, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 105, 1005}

	v := slices.Clone(in)
	DoubleCutoff(v, CUTOFF)
	if !reflect.DeepEqual(v, exp) {
		t.Fatal("DoubleCutoff (scalar) failed, got:", v)
	}

	v = slices.Clone(in)
	ok := DoubleCutoffSIMD(v, 5)
	if !ok {
		t.Log("WARNING: DoubleCutoffSIMD did not use SIMD")
	}
	if !reflect.DeepEqual(v, exp) {
		t.Fatal("DoubleCutoffSIMD (vector) failed, got:", v)
	}
}

// DoubleCutoff doubles the values if less than "cutoff" else the cutoff is just added
func DoubleCutoff(a []int16, cutoff int) { // 24ns
	c := int16(cutoff)
	for i := 0; i < len(a); i++ {
		if a[i] < c {
			a[i] *= 2
		} else {
			a[i] += c
		}
	}
}

func DoubleCutoffSIMD(a []int16, cutoff int) bool {
	in := archsimd.LoadInt16x16Slice(a)            // VMOVDQU
	c := archsimd.BroadcastInt16x16(int16(cutoff)) // VBROADCASTW

	mask := in.Less(c).ToInt16x16() // VPCMPGTW
	double := in.Add(in).And(mask)  // VPADDW, VPAND
	add := in.Add(c).AndNot(mask)   // VPADDW, VPANDN

	// Or the 2 vectors to get the result and return it
	double.Or(add).StoreSlice(a) // VPOR, VPMOVDQU
	return true
}

func DoubleCutoffSIMD2(a []int16, cutoff int) bool {
	in := archsimd.LoadInt16x16Slice(a)
	if !archsimd.X86.AVX2() || len(a) != in.Len() {
		DoubleCutoff(a, cutoff)
		return false
	}

	c := archsimd.BroadcastInt16x16(int16(cutoff))
	mask := in.Less(c).ToInt16x16()                               // which elements are less than cutoff
	in.Add(in).And(mask).Or(c.Add(in).AndNot(mask)).StoreSlice(a) // OR the 2 vectors and store the result
	return true
}

func BenchmarkDoubleCutoff(b *testing.B) { // 24ns
	for b.Loop() {
		DoubleCutoff([]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 100, 1000}, 5)
	}
}

func BenchmarkDoubleCutoffSIMD(b *testing.B) { // 62ns
	for b.Loop() {
		DoubleCutoffSIMD([]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 100, 1000}, 5)
	}
}

func DoubleCutoffSIMDArray(a *[16]int16, cutoff int) bool {
	in := archsimd.LoadInt16x16(a) // this line makes it slow (55ns/op) vs next line (3 ns/op)
	//in := archsimd.LoadInt16x16(&[16]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 100, 1000}) // FAST
	c := archsimd.BroadcastInt16x16(int16(cutoff))

	mask := in.Less(c).ToInt16x16()
	double := in.Add(in).And(mask)
	add := in.Add(c).AndNot(mask)

	// Or the 2 vectors to get the result and return it
	double.Or(add).Store(a)
	return true
}

func BenchmarkDoubleCutoffSIMDArray(b *testing.B) { // 62ns
	for b.Loop() {
		DoubleCutoffSIMDArray(&[16]int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 100, 1000}, 5)
	}
}

//////////////////////////////////////////////////////

func DCPartSlowOrig() int16 { // 63 ns/op
	a := make([]int16, 16)
	var in archsimd.Int16x16
	in = archsimd.LoadInt16x16Slice(a) // **** remove this to make fast ****
	c := archsimd.BroadcastInt16x16(int16(5))

	mask := in.Less(c).ToInt16x16()
	double := in.Add(in).And(mask)
	add := in.Add(c).AndNot(mask)

	double.Or(add).StoreSlice(a)
	return a[0]
}

func DCPartSlow2() int16 { // 58 ns/op
	a := make([]int16, 16)
	var in archsimd.Int16x16
	in = archsimd.LoadInt16x16Slice(a)
	c := archsimd.BroadcastInt16x16(int16(5))

	double := in.Add(in)
	add := in.Add(c)

	double.Or(add).StoreSlice(a)
	return a[0]
}

func DCPartSlow3() int16 { // 54 ns/op
	a := make([]int16, 16)
	var in archsimd.Int16x16
	in = archsimd.LoadInt16x16Slice(a)
	c := archsimd.BroadcastInt16x16(int16(5))

	v := in.Add(c)

	v.Or(v).StoreSlice(a)
	return a[0]
}

func DCPartSlow4() int16 { // 54 ns/op
	a := make([]int16, 16)
	var in archsimd.Int16x16
	in = archsimd.LoadInt16x16Slice(a)

	in.Add(in).StoreSlice(a)
	return a[0]
}

func DCPartSlow5() int { // 10 ns/op
	var in archsimd.Int16x16
	in = archsimd.LoadInt16x16SlicePart([]int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	in = in.Add(in)
	return 1
}

func DCPartSlow() int { // 1.67
	var in archsimd.Int16x16
	in = archsimd.LoadInt16x16(&[16]int16{1})

	in = in.Add(in)
	return 1
}

func DCPartFastOrig() int16 { // 2.2ns/op
	var in archsimd.Int16x16
	c := archsimd.BroadcastInt16x16(int16(5))

	mask := in.Less(c).ToInt16x16()
	double := in.Add(in).And(mask)
	add := in.Add(c).AndNot(mask)

	a := make([]int16, 16)
	double.Or(add).StoreSlice(a)
	return a[0]
}

func DCPartFast() int16 { // 2.2ns/op
	//in := archsimd.LoadInt16x16SlicePart([]int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6})  // *** Slow with this line ***
	//in := archsimd.LoadInt16x16SlicePart([]int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6}[12:])
	in := archsimd.LoadInt16x16(&[16]int16{}) // *** FAST ***
	c := archsimd.BroadcastInt16x16(int16(5))

	mask := in.Less(c).ToInt16x16()
	double := in.Add(in).And(mask)
	add := in.Add(c).AndNot(mask)

	a := make([]int16, 16)
	double.Or(add)
	return a[0]
}

func BenchmarkDoubleCutoffSIMDOverhead(b *testing.B) {
	for b.Loop() {
		DCPartFast()
	}
}
