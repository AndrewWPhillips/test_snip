package __

import (
	"testing"
)

type Seq0 func(yield func() bool) // like iter.Seq and iter.Seq2 but yields zero values

// ZeroRange is a range func yielding no values
func ZeroRange(max int) Seq0 {
	return func(yield func() bool) {
		for i := 0; i < max; i++ {
			if !yield() {
				return
			}
		}
	}
}

func TestSeq0(t *testing.T) {
	for range ZeroRange(5) {
		println("X")
	}
}
