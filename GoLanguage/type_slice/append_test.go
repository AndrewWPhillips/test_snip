package __

// tests of changing slices by reslice and append - see also extend_test

import (
	"log"
	"testing"
)

func TestAppend(t *testing.T) {
	// you can increase slice length past cap but this may require realloc - returned slice may have different address to passed in one
	tmp := []int{1, 2, 3, 4, 5}
	si := append(tmp, 6, 7)
	log.Printf("%p %p %d %d", tmp, si, len(si), cap(si)) // 7 10 - cap is twice previous size (only guaranteed to be at least 7)
}

func TestAppend2(t *testing.T) {
	si := make([]int, 0, 5)
	log.Printf("len %d cap %d\n", len(si), cap(si)) // 0 5
	si = si[1:2]
	log.Printf("%p len %d cap %d\n", &si[0], len(si), cap(si)) // 1 4
	si = si[1:2]
	log.Printf("%p len %d cap %d\n", &si[0], len(si), cap(si)) // 1 3
	s2 := append(si, 42, 47, 71, 73)
	log.Printf("%p len %d cap %d\n", &si[0], len(si), cap(si)) // 1 3 (si not changed)
	log.Printf("%p len %d cap %d\n", &s2[0], len(s2), cap(s2)) // 5 6 (s2 has different address to si)
	si = append(si, 1, 2)
	log.Printf("%p len %d cap %d\n", &si[0], len(si), cap(si)) // 3 3 (new length but same address)
}

func TestExpand(t *testing.T) {
	s := make([]string, 2, 4)
	s[0] = "zero"
	s[1] = "one"

	log.Printf("%p len %d cap %d\n", &s[0], len(s), cap(s))
	s = append(s, "two")
	log.Printf("%p len %d cap %d\n", &s[0], len(s), cap(s)) // len 3 cap 4
	s = append(s, "three", "four")
	log.Printf("%p len %d cap %d\n", &s[0], len(s), cap(s)) // len 5 cap 8 (s[0] has diff address)
}

func TestAppendNil(t *testing.T) {
	var a, b []string
	a = append(a, b...) // append nil slice to nil slice
	println(a == nil, len(a))
}
// TestAppendNilCap tests cap after creating a new slice by appending to a new slice
// Note that this behaviour is unspecified so may change at some point
func TestAppendNilCap(t *testing.T) {
	var a []int
	a = append(a, 1)
	println(cap(a)) // 1
	a = append(a, 2)
	println(cap(a)) // 2
	a = append(a, 3)
	println(cap(a)) // 4
	a = append(a, 4)
	println(cap(a)) // 4
	a = append(a, 5)
	println(cap(a)) // 8 - could be less in the future?? (cf C++ std::vector)
	a = append(a, 6)
	println(cap(a)) // 8
}
