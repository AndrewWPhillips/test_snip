package __

import (
	"testing"
)

func BenchmarkDeclare(b *testing.B) { // 0 allocs/op
	var outer []int
	for i := 0; i < b.N; i++ {
		var s []int
		outer = s
	}
	TestGlobal = outer
}

func BenchmarkLiteral(b *testing.B) { // 1 allocs/op
	var outer []int
	for i := 0; i < b.N; i++ {
		s := []int{1, 2, 3, 4}
		outer = s
	}
	TestGlobal = outer
}

func BenchmarkMake(b *testing.B) { // 1 allocs/op
	var outer []int
	for i := 0; i < b.N; i++ {
		s := make([]int, 4)
		outer = s
	}
	TestGlobal = outer
}

func BenchmarkArray(b *testing.B) { // 0 allocs/op
	var a [4]int
	var outer []int
	for i := 0; i < b.N; i++ {
		s := a[1:2]
		outer = s
	}
	TestGlobal = outer
}
