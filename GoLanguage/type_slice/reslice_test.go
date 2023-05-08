package __

import (
	"fmt"
	"log"
	"testing"
)

func TestResliceBeyondBounds(t *testing.T) {
	s := "0123456789ABC"
	log.Println(s[:99]) // panic: runtime error: slice bounds out of range
}

// TestResliceLiterals demonstrates how to slice an array literal
func TestResliceLiterals(t *testing.T) {
	//log.Println([4]int{42, 43, 44}[:2]) // slice of unaddressable value

	log.Println((&[4]int{42, 43, 44})[:2]) // OK using address of array

	a := [4]int{42, 43, 44}
	log.Println(a[:2]) // OK if assign to a variable first

	log.Println([]int{42, 43, 44, 0, 0}[:2]) // OK can re-slice a slice literal

	log.Println("23400"[:2]) // OK to re-slice a string literal
}

func TestStartAfterEnd(t *testing.T) {
	v := [...]int{0, 1, 2, 3, 4, 5}
	w := v[:2]
	fmt.Println(w[3:5])
	fmt.Println(w[3:])
}
