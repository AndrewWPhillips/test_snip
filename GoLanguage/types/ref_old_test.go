package __

// These are tests whether different types are "reference" type in Go. The Go language actually does not mention ref.
// type (anymore) but many people say that at least map and slice are ref. types meaning that if you create a new one
// by assigning an old one then modify the new one you also modify the old one.  This is true for map and chan,
// partially true for slices, but not true for immutable types like interface, slice and func.
// Of course, pointers, by definition, allow you to modify what is pointed to.

import (
	"log"
	"testing"
)

// TestIsMapRefType shows that copying a map only copied the "pointer" to the data
func TestIsMapRefType(t *testing.T) {
	v := map[int]string{1: "one"}
	ref := v
	ref[1] = "42"  // changes contents of the original map
	log.Println(v) // map[1:42]
}

// TestIsSliceRefType shows that copying a slice copies "pointer" to the data, but not len and cap
func TestIsSliceReferenceType(t *testing.T) {
	v := []int{0, 1, 2, 3, 4}
	v = v[0:3] // len=3, cap=5
	ref := v
	ref[1] = 42    // change contents of original slice via the new slice
	log.Println(v) // [0 42 2]

	//ref = ref[::3] // ERROR (why?): 2nd index required in 3-index slice
	ref = ref[:2:3] // change capacity of v
	log.Println(v, ref)
	log.Println(cap(v), cap(ref))
}

// TestIsChanRefType shows that copying a chan only copied the "pointer"
func TestIsChanReferenceType(t *testing.T) {
	v := make(chan int, 2)
	ref := v
	ref <- 42
	v <- 1
	log.Println(<-v) // 42
}

// TestIsFuncRefType shows that a func is not a ref type
func TestIsFuncReferenceType(t *testing.T) {
	v := func() int { return 1 }
	ref := v
	ref = func() int { return 2 }
	log.Println(v())
	log.Println(ref())

	// You can take the address of a func and modify it through that
	ptr := &v
	*ptr = func() int { return 3 }
	log.Println(v())
}

// TestIsPtrRefType shows that copying a pointer allows you to change the original
func TestIsPtrRefType(t *testing.T) {
	a := 1
	v := &a
	ref := v
	*ref = 42
	log.Println(*v)
}
