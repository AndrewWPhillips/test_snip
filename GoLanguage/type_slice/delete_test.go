package __

import (
	"log"
	"testing"
)

// TestDeleteNoOrder tests deleting an element without preserving order (last elt is moved to the gap)
func TestDeleteNoOrder(t *testing.T) {
	a := (&[10]int{1, 2, 3, 4, 5, 6})[:6]
	a[3] = a[len(a)-1]
	a = a[:len(a)-1]
	log.Println(a)
}

// TestDeleteKeepOrder tests deleting from the slice while preserving order
func TestDeleteKeepOrder(t *testing.T) {
	a := (&[10]int{1, 2, 3, 4, 5, 6})[:6]
	copy(a[3:], a[5:]) // delete indices 3 and 4
	a = a[:len(a)-2]

	//a = append(a[:3], a[5:]...)

	//a = a[:3+copy(a[3:], a[5:])]

	log.Println(a)
}

// TestDeleteRetain creates a slice with elements deleted, retaining the old slice
func TestDeleteRetain(t *testing.T) {
	// create new slice with 2 elements [3:5] deleted
	a := (&[10]int{1, 2, 3, 4, 5, 6})[:6]
	b := append([]int{}, a[:3]...)
	b = append(b, a[5:]...)

	log.Println(a)
	log.Println(b)
}

func TestSlicingArray(t *testing.T) {
	//s := [10]int{1, 2, 3, 4, 5, 6}[1:2] // why is this illegal?
	//s := (&[10]int{1, 2, 3, 4, 5, 6})[1:2] // OK
	//s := "abc"[1:2] // OK
	a := [10]int{1, 2, 3, 4, 5, 6}
	z := a[1:4]
	log.Println(z)
}
