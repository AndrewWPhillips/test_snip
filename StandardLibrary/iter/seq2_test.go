package __

import (
	"iter"
	"testing"
)

func TwoIntoOne[T any](in iter.Seq[T]) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		var saved T
		doIt := false // F-T-F... so we only yield on every other iteration
		for i := range in {
			if doIt {
				if !yield(saved, i) {
					return
				}
			} else {
				saved = i
			}
			doIt = !doIt
		}
		if doIt {
			// If there is one left send it and a "zero" value
			yield(saved, *new(T))
		}
	}
}

func TestTwoIntoOne(t *testing.T) {
	for i, j := range TwoIntoOne(OneRange(5)) {
		println(i, j)
	}
}
