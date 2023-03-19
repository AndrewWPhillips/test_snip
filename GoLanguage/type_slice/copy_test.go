package __

import (
	"log"
	"testing"
)

func TestCopy(t *testing.T) {
	a := (&[10]int{1, 2, 3, 4, 5})[:3]
	b := append([]int{}, a...)

	c := make([]int, len(a)-2, 20)
	copy(c, a)
	d := make([]int, len(a), 20)
	copy(d, a)
	e := make([]int, len(a)+2, 20)
	copy(e, a)

	log.Printf("%d %v\n", len(a), a)
	log.Printf("%d %v\n", len(b), b)
	log.Printf("%d %v\n", len(c), c)
	log.Printf("%d %v\n", len(d), d)
	log.Printf("%d %v\n", len(e), e)
}

func TestSliceCopy(t *testing.T) {
	a := []byte{1, 2, 3}
	b := a
	c := a[:]
	d := make([]byte, len(a))
	copy(d, a)
	a[0] = 42
	log.Println(b) // [42 2 3] changing a changes b
	log.Println(c) // [42 2 3] changing a changes c
	log.Println(d) // [1 2 3] OK
}
