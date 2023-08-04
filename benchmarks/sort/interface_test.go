package __

import (
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkOldSort0(b *testing.B) { // 101ns/op
	s := make([]int, SliceLength)
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}

func BenchmarkOldSort1(b *testing.B) { // 106ns/op
	j := 0
	s := make([]int, SliceLength)
	for n := range s {
		s[n] = rand.Int()
	}
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
		j = s[0]
	}
	println(j)
}

// BenchmarkOldSort2 checks sorting slice each time with a new
func BenchmarkOldSort2(b *testing.B) { // 302ns/op
	j := 0
	src := make([]int, SliceLength)
	for n := range src {
		src[n] = rand.Int()
	}
	s := make([]int, SliceLength)
	for i := 0; i < b.N; i++ {
		copy(s, src) // ~4ns/op see BenchmarkSliceCopy
		sort.Ints(s)
		j = s[0]
	}
	println(j)
}
