package main

import (
	"iter"
	"slices"
)

func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

var ss = []string{
	"jlfgjlfdkjgflkd",
	"a",
	"dsaldj",
	"3",
	"4",
	"5",
}

var global string

func main() {
	for _, v := range slices.Backward(ss) {
		global = v
	}
}
