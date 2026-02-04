package main

import "iter"

func Backward2[Slice ~[]E, E any](s Slice) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s); i > 0; i-- {
			if !yield(i-1, s[i-1]) {
				return
			}
		}
	}
}

var ss2 = []string{
	"jlfgjlfdkjgflkd",
	"a",
	"dsaldj",
	"3",
	"4",
	"5",
}

func main() {
	for i, v := range Backward2(ss2) {
		println(i, v)
	}
}
