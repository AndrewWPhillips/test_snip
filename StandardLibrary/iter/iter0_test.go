//go:build go1.23

package __

import (
	"testing"
)

func TestIter0(t *testing.T) {
	var i int
	f := func(yield func() bool) {
		for i < 10 {
			if !yield() {
				return
			}
			i++
		}
	}

	for range f {
		println("X")
	}
}

var i int

func It(yield func(int) bool) {
	for i < 10 {
		if !yield(i) {
			return
		}
		i++
	}
}

func TestIter0A(t *testing.T) {
	for n := range It {
		println(n)
	}
}

func TestIter0B(t *testing.T) {
	for n := range It {
		println(n)
	}
}

type iter_t int

func (it iter_t) Iterate(yield func(int) bool) {
	for i < 10 {
		if !yield(i) {
			return
		}
		i++
	}
}

func TestIter0C(t *testing.T) {
	var it iter_t
	for n := range it.Iterate {
		println(n)
	}
}
