package __

import (
	"math/rand"
	"testing"

	"golang.org/x/exp/slices"
)

func BenchmarkGenericSort0(b *testing.B) { // ~22ns/op (work computer)
	s := []int{1, 2, 3, 42, 77, 99, 47, 72, 1, 2, 3, 42, 77, 99, 47, 72}
	for i := 0; i < b.N; i++ {
		slices.Sort(s)
	}
}

func BenchmarkGenericSort1(b *testing.B) { // ~22ns/op (work computer)
	j := 0
	s := make([]int, SliceLength)
	for n := range s {
		s[n] = rand.Int()
	}
	for i := 0; i < b.N; i++ {
		slices.Sort(s)
		j = s[0]
	}
	println(j)
}

const SliceLength = 16

// BenchmarkSliceCopy just checks time for slice copy, so we can subtract from below
// to work out the sort time without the slice copying time
func BenchmarkSliceCopy(b *testing.B) { // ~4ns/op with SliceLength=16 (work computer)
	j := 0
	src := make([]int, SliceLength)
	for n := range src {
		src[n] = rand.Int()
	}
	s := make([]int, SliceLength)
	for i := 0; i < b.N; i++ {
		copy(s, src)
		j = s[0]
	}
	println(j)
}

// BenchmarkGenericSort2 checks sorting slice each time with a new
func BenchmarkGenericSort2(b *testing.B) { // 68ns/op  (68-4 => ~64 for sort)
	j := 0
	src := make([]int, SliceLength)
	for n := range src {
		src[n] = rand.Int()
	}
	s := make([]int, SliceLength)
	for i := 0; i < b.N; i++ {
		copy(s, src)
		slices.Sort(s)
		j = s[0]
	}
	println(j)
}

// BenchmarkSliceRandomise checks time for setting the slice so we can subtract from below
// to work out the sort time without the slice w/o slice init time
func BenchmarkSliceRandomise(b *testing.B) { // 93 ns/op
	s := make([]int, SliceLength)
	for i := 0; i < b.N; i++ {
		for n := range s {
			s[n] = rand.Int()
		}
	}
}

// BenchmarkGenericSort3 checks sorting slice each time with a new random slice
func BenchmarkGenericSort3(b *testing.B) { // 315ns/op (315-93 => ~220ns/op for sort)
	s := make([]int, SliceLength)
	for i := 0; i < b.N; i++ {
		for n := range s {
			s[n] = rand.Int()
		}
		slices.Sort(s)
	}
}

// BenchmarkGenericSort4 = same but exclude slice init from times
// NOTE: this seems to take too long - prob stop/start in the loop confusing the timings
func BenchmarkGenericSort4(b *testing.B) {
	j := 0
	s := make([]int, SliceLength)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for n := range s {
			s[n] = rand.Int()
		}
		b.StartTimer()
		slices.Sort(s)
		j = s[0]
	}
	println(j)
}
