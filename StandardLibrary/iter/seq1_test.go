package __

import (
	"fmt"
	"iter"
	"strconv"
	"testing"
)

// OneRange is a range func yielding one (int) value
func OneRange(max int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < max; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// SliceRange is a range func yielding one (string slice) value
// This shows that the type of the yield value can be anything.
func SliceRange(max int) iter.Seq[[]string] {
	return func(yield func([]string) bool) {
		for i := 0; i < max; i++ {
			r := make([]string, 0, i)
			for n := range i {
				r = append(r, strconv.Itoa(n))
			}
			if !yield(r) {
				return
			}
		}
	}
}

func TestSeq1(t *testing.T) {
	for i := range OneRange(5) {
		println(i)
	}
}

func TestSeq1Slice(t *testing.T) {
	for v := range SliceRange(5) {
		fmt.Println(v)
	}
}
