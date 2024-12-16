package __

import (
	"iter"
	"testing"
)

func TestPullOne(t *testing.T) {
	next, stop := iter.Pull(OneRange(5))
	defer stop()

	println(next())
	println(next())
	println(next())
	println(next())
	println(next())
	println(next())
}

func TestPullVsRange(t *testing.T) {
	for i := range OneRange(5) {
		println(i)
	}

	next, stop := iter.Pull(OneRange(5))
	defer stop()
	for i, ok := next(); ok; i, ok = next() {
		println(i)
	}
}

// TwoToOne is like TwoIntoOne but uses a pull iterator
func TwoToOne[T any](in iter.Seq[T]) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		// Get a pull iterator from in
		next, stop := iter.Pull(in)
		defer stop()
		for {
			if a, ok := next(); ok {
				b, _ := next()
				if !yield(a, b) {
					return
				}
			}
		}
	}
}

func Test221Pull(t *testing.T) {
	for i, j := range TwoToOne(OneRange(5)) {
		println(i, j)
	}
}
