package __

import (
	"testing"
)

func BenchmarkStringNoAllocs(b *testing.B) { // BenchmarkStringNoAllocs-12  1000000000  0.2628 ns/op  0 B/op 0 allocs/op
	b.ReportAllocs()
	outer := ""
	for i := 0; i < b.N; i++ {
		s := ""
		outer = s
	}
	println(outer)
}

func BenchmarkStringAllocs(b *testing.B) { // BenchmarkStringAllocs-12   71582815   16.55 ns/op  4 B/op  1 allocs/op
	b.ReportAllocs()
	outer := ""
	for i := 0; i < b.N; i++ {
		s := string(rune(i))
		outer = s
	}
	println(outer)
}

func BenchmarkStringRuneArray(b *testing.B) { // same as BenchmarkStringAllocs above
	b.ReportAllocs()
	outer := ""
	a := [4]rune{}
	for i := 0; i < b.N; i++ {
		outer = string(a[i%4])
	}
	println(outer)
}

func BenchmarkStringArray(b *testing.B) { // 1000000000  0.7978 ns/op  0 B/op 0 allocs/op
	b.ReportAllocs()
	outer := ""
	a := [4]string{string(rune(0)), string(rune(1)), string(rune(2)), string(rune(3))}
	for i := 0; i < b.N; i++ {
		outer = a[i%4]
	}
	println(outer)
}
