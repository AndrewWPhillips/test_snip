package __

import (
	"iter"
	"testing"
)

var ss = []string{
	"jlfgjlfdkjgflkd",
	"a",
	"dsaldj",
	"3",
	"4",
	"5",
}

// Backward returns an iterator over index-value pairs in the slice,
// traversing it backward with descending indices.
func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func BenchmarkSliceReverse(b *testing.B) { // best: 151 ns/op
	//for b.Loop() {
	//	for _, v := range slices.Backward(ss) {
	//		_ = v
	//	}
	//}
	for b.Loop() {
		for range Backward(ss) {
		}
	}
}

func Backward2[Slice ~[]E, E any](s Slice) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s); i > 0; i-- {
			if !yield(i-1, s[i-1]) {
				return
			}
		}
	}
}

func BenchmarkMyReverse(b *testing.B) { // best = 137 ns/op
	for b.Loop() {
		for range Backward2(ss) {
		}
	}
}
