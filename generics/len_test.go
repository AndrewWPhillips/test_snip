package __

import (
	"testing"
)

/*
func Length[U any, V comparable, W any, T []U | map[V]W](v T) int {
	return len(v)
}

func TestLength(t *testing.T) {
	println(Length[int, int, int, []int]([]int{1, 2, 3}))
}
*/

type SliceOrMapOfSlices[K, KK comparable, V any] interface {
	~map[K]map[KK]V | ~map[K][]V
}

func CountElements[K, KK comparable, V any, T SliceOrMapOfSlices[K, KK, V]](raw T) int {
	num := 0
	// This gives ERROR: cannot range over raw (variable of type T constrained by SliceOrMapOfSlices[K, KK, V]): no core type
	// unless we remove one of the above 2 types from SliceOrMapOfSlices
	//for _, elements := range raw {
	//	num += len(elements)
	//}
	return num
}

func TestCount(t *testing.T) {
	m := map[int]map[int]int{
		1: map[int]int{0: 1, 1: 2, 2: 3},
		5: map[int]int{0: 4, 1: 5},
	}
	println(CountElements[int, int, int, map[int]map[int]int](m))
}
