package __

import (
	"log"
	"testing"
)

// TestPop "removes" element 0 from the slice
// Note that no memory is allocated/freed - backing array is unchanged
func TestPop(t *testing.T) {
	a := (&[10]int{1, 2, 3, 4, 5, 6})[:6]
	a = a[1:]
	log.Println(a)
}

// TestPush adds an element to the front of a slice
// Note that a new backing array has to be allocated (and old a may be deleted if nothing else uses it)
func TestPush(t *testing.T) {
	a := (&[10]int{1, 2, 3, 4, 5, 6})[:6]
	a = append([]int{7}, a...)
	log.Println(a)
}
