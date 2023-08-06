package __

// These are tests whether different types are "reference" type in Go. The Go language actually does not mention ref.
// type (anymore) but many people say that at least map and slice are ref. types meaning that if you create a new one
// by assigning an old one then modify the new one you also modify the old one.  This is true for map and chan,
// partially true for slices, but not true for immutable types like interface and string.
// Of course, pointers, by definition, allow you to modify what is pointed to.

// For clarity, I define a reference type T as having the behaviour of the following steps:
// 1. var a, b T
// 2. // assign/modify a somehow (indirection, element, field, etc)
// 4. b = a
// 5. // modify b somehow
// 6. println(a)   // changes in b are reflected in a

import (
	"log"
	"testing"
)

// TestRefPointer shows that copying a pointer allows you to change the original
func TestRefPointer(t *testing.T) {
	var a, b *int
	n := 1
	a = &n
	b = a
	*b = 42
	log.Println(*a) // 42
}

type s1 struct {
	n int
}

// TestRefStruct1 shows that a struct is *not* normally a "reference" type
func TestRefStruct1(t *testing.T) {
	var a, b s1
	a.n = 1
	b = a
	b.n = 42
	log.Println(a.n) // 1
}

type s2 struct {
	s []int
}

// TestRefStruct2 shows that a struct is a "reference" type if it contains a reference type
func TestRefStruct2(t *testing.T) {
	var a, b s2
	a.s = make([]int, 1)
	b = a
	b.s[0] = 42
	log.Println(a.s) // [42]
}

// TestRefInterface1 shows that an interface is a *not* normally a "reference"
func TestRefInterface1(t *testing.T) {
	var a, b interface{}
	a = 1
	b = a
	b = 42
	log.Println(a, b) // 1 42
}

// TestRefInterface2 shows that an interface is a "reference" type if it contains a reference type
func TestRefInterface2(t *testing.T) {
	var a, b interface{}
	a = make([]int, 1)
	b = a
	b.([]int)[0] = 42
	log.Println(a.([]int)) // [42]
}

// TestRefMap shows that copying a map only copied the "pointer" to the data
func TestRefMap(t *testing.T) {
	var a, b map[int]string
	a = map[int]string{1: "one"}
	b = a
	b[1] = "42"    // changes contents of the original map
	log.Println(a) // mp[1:42]
}

func TestRefSlice(t *testing.T) {
	var a, b []int
	a = []int{0, 1, 2}
	b = a
	b[1] = 42
	log.Println(a) // [0 42 2]
}

// TestRefChan shows that copying a chan only copies the "pointer"
func TestIsChanRefType(t *testing.T) {
	var a, b chan int
	a = make(chan int, 2)
	b = a
	b <- 42
	log.Println(<-a) // 42
}

type myInt int

func (m *myInt) f() int {
	*m++
	return int(*m)
}

// TestRefFunc shows that a func is a reference type
func TestRefFunc(t *testing.T) {
	var a, b func() int
	m := myInt(1)
	a = m.f
	b = a
	b()
	log.Println(a()) // 42
}

// TestIsSliceRefType shows that copying a slice copies "pointer" to the data, but not len and cap
func TestIsSliceRefType(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4}
	v := a[0:2] // len=2, cap=5
	ref := v
	ref[1] = 42       // change contents of original slice via the new slice
	log.Println(v[1]) // 42

	//ref = ref[::3] // ERROR (why?): 2nd index required in 3-index slice
	ref = ref[:2:3] // change capacity of v
	log.Println(v, ref)
	log.Println(cap(v), cap(ref))
}

// TestIsFuncRefType shows that a func is not a ref type
func TestIsFuncRefType(t *testing.T) {
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

type sp struct {
	p *int
}

func TestStructRefType(t *testing.T) {
	n := 1
	a := sp{p: &n}
	b := a
	*b.p = 42
	log.Println(*a.p) // 42
}
