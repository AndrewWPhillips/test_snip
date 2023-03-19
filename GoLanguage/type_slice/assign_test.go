package __

// Tests of creating a slice using make, or slicing an array or slice

import (
	"log"
	"testing"
)

var a = [...]int{1, 2, 3, 4, 5}

func TestMake(t *testing.T) {
	var si []int

	si = make([]int, 0)
	log.Printf("%v %p %d %d\n", si, si, len(si), cap(si))
	si = make([]int, 0, 4)
	log.Printf("%v %p %d %d\n", si, si, len(si), cap(si))
	si = make([]int, 2, 4)
	log.Printf("%v %p %d %d\n", si, si, len(si), cap(si))
	//si = make([]int, 2, 0) // ERROR: invalid argument: length and capacity swapped
}

func TestFromArray(t *testing.T) {
	var si []int

	si = a[:]
	log.Printf("%v %p %d %d\n", si, si, len(si), cap(si))
	si = a[1:3]
	log.Printf("%v %p %d %d\n", si, si, len(si), cap(si))
	//si = [4]int{}[:2] // error: slice of unaddressable value
	tmp := [4]int{}
	si = tmp[:2] // SIMPLER: si = make([]int, 2, 4)
	log.Printf("%v %p %d %d\n", si, si, len(si), cap(si))
}

func TestReslice(t *testing.T) {
	var si []int

	si = []int{1, 2, 3, 4, 5}[1:2]
	log.Printf("%v %d %d\n", si, len(si), cap(si)) // [2] 1 4

	si = si[:3]                                         // can increase slice length using up to cap
	log.Printf("increase len: %d %d", len(si), cap(si)) // 3 4

	//si = []int{1, 2, 3, 4, 5}[:99] // runtime error: slice bounds out of range

	// you can't reslice backwards since a slice only stores len and cap (no cap back from start)
	//si = []int{1, 2, 3, 4, 5}[2:];	si = si[-2:] // compile error: invalid argument: index -2 (constant of type int) must not be negative

	si = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}[1:2:5]             // can reduce cap
	log.Printf("reduce cap: %v %d %d", si, len(si), cap(si)) // [2] 1 4

	// you can increase slice length past cap but this may require realloc - returned slice may have different address to passed in one
	tmp := []int{1, 2, 3, 4, 5}
	si = append(tmp, 6, 7)
	log.Printf("%p %p %d %d", tmp, si, len(si), cap(si)) // 7 10 - cap is twice previous size (only guaranteed to be at least 7)
}
