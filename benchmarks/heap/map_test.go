package __

import (
	"strconv"
	"testing"
)

func BenchmarkEmptyAllocs(b *testing.B) { // 1 alloc/op
	b.ReportAllocs()
	var outer map[int]string
	for i := 0; i < b.N; i++ {
		m := make(map[int]string)
		outer = m
	}
	TestGlobal = outer
}

func BenchmarkStringAllocs(b *testing.B) { // 2 alloc/op
	b.ReportAllocs()
	var outer map[int]string
	for i := 0; i < b.N; i++ {
		m := map[int]string{i: ""}
		outer = m
	}
	TestGlobal = outer
}

func BenchmarkString2Allocs(b *testing.B) { // 2 allocs
	b.ReportAllocs()
	var outer map[int]string
	for i := 0; i < b.N; i++ {
		m := map[int]string{i % 4: "", i: strconv.Itoa(i % 16)}
		outer = m
	}
	TestGlobal = outer
}

func BenchmarkString3Allocs(b *testing.B) { // 2 allocs
	b.ReportAllocs()
	var outer map[int]string
	for i := 0; i < b.N; i++ {
		m := map[int]string{i % 4: "a", i%5 + 5: "b", i%2 + 10: "c"}
		outer = m
	}
	TestGlobal = outer
}

func BenchmarkIntAllocs(b *testing.B) { // 2 alloc/op
	b.ReportAllocs()
	var outer map[int]int
	for i := 0; i < b.N; i++ {
		m := map[int]int{i % 4: i}
		outer = m
	}
	TestGlobal = outer
}

func BenchmarkSliceIntAllocs(b *testing.B) { // 1 alloc
	b.ReportAllocs()
	var outer []int
	for i := 0; i < b.N; i++ {
		s := make([]int, 4)
		s[i%4] = i
		outer = s
	}
	TestGlobal = outer
}
