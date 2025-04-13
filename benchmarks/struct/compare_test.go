package __

import (
	"testing"
)

func BenchmarkCompareInt(b *testing.B) { // 0.28 ns/op
	j, k, r := 2, 3, false
	for i := 0; i < b.N; i++ {
		r = j == k
	}
	println(r)
}

func BenchmarkCompareIntNew(b *testing.B) { // 2.2 ns/op ???
	i, j := 2, 2
	for b.Loop() {
		_ = i == j
	}
}

type (
	b8 struct {
		a int32
		b bool
	}

	b8a struct {
		a [8]byte
		b [7]byte
	}

	b16 struct {
		a int64
		b bool
	}
)

func BenchmarkCompareStruct8(b *testing.B) { // 0.8 ns/op
	var j, k b8a
	var r bool
	for i := 0; i < b.N; i++ {
		r = j == k
	}
	println(r)
}
func BenchmarkCompareStruct8Diff(b *testing.B) { // 0.28 ns/op
	j := b8{a: 1, b: false}
	k := b8{a: 2, b: true}
	var r bool
	for i := 0; i < b.N; i++ {
		r = j == k
	}
	println(r)
}

func BenchmarkCompareStruct8Diff2(b *testing.B) { // 0.28 ns/op
	j := b8{a: 1, b: false}
	k := b8{a: 1, b: true}
	var r bool
	for i := 0; i < b.N; i++ {
		r = j == k
	}
	println(r)
}
func BenchmarkCompareStruct8Diff3(b *testing.B) {
	j := b8{a: 0, b: false}
	k := b8{a: 0, b: false}
	var r bool
	for i := 0; i < b.N; i++ {
		r = j == k
	}
	println(r)
}

func BenchmarkCompareStruct8a(b *testing.B) { // 0.84 ns/op
	var j, k b8a
	var r bool
	for i := 0; i < b.N; i++ {
		r = j == k
	}
	println(r)
}

func BenchmarkCompareStruct16(b *testing.B) { // 0.29 ns/op
	j := b16{a: 1, b: false}
	k := b16{a: 2, b: true}
	var r bool
	for i := 0; i < b.N; i++ {
		r = j == k
	}
	println(r)
}
