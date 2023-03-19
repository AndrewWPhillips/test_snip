package __

import (
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	a := (&[10]int{1, 2, 3, 4, 5, 6})[:6]
	//a = append(append(a[:2], []int{42,43,44}...), a[2:]...)  // wrong!!!
	a = append(a[:2], append([]int{42, 43, 44}, a[2:]...)...) // inefficient!
	log.Println(a)
}

// TestInsertRetain tests inserting into a slice while retaining the original
// Note efficiency could be improved since it may cause 3 heap allocations when only 1 is necessary
func TestInsertRetain(t *testing.T) {
	a := (&[10]int{1, 2, 3, 4, 5, 6})[:6]
	b := append([]int{}, a[:2]...)
	b = append(b, []int{42, 43, 44}...)
	b = append(b, a[2:]...)
	log.Println(a)
	log.Println(b)
}
