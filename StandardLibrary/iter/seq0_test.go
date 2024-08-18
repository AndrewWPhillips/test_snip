package __

import (
	"testing"
)

type Seq0 func(yield func() bool)

func All() Seq0 {
	i := 0
	return func(yield func() bool) {
		for i < 10 {
			if !yield() {
				return
			}
			i++
		}
	}
}

func TestSeq0(t *testing.T) {
	for range All() {
		println("X")
	}
}
