package __

import (
	"log"
	"testing"

	"golang.org/x/exp/slices"
)

// TestSliceEqual1 shows that slices compare equal even if their capacities are different
func TestSliceEqual1(t *testing.T) {
	a := make([]int, 3, 5)
	b := make([]int, 3, 6)
	a[1], b[1] = 42, 42
	log.Println(slices.Equal(a, b))
}

// TestSliceEqualDepth tests how "deep" the equality test of slices.Equal is
// As expected it only compares each element using a shallow comparison
func TestSliceEqualDepth(t *testing.T) {
	i, j := 42, 42
	a := []pointerStruct{{p: &i}}
	b := []pointerStruct{{p: &j}}
	c := []pointerStruct{{p: &j}}
	log.Println(slices.Equal(a, b)) // false
	log.Println(slices.Equal(b, c)) // true
}

func TestSlicesReplace(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := slices.Replace(a, 1, 6, 22, 33)
	log.Println(a)
	log.Println(b)
}
