package __

import (
	"log"
	"testing"
)

// Puzzle from Brad Fitzpatrick https://twitter.com/bradfitz/status/951534740405829632
func TestBradsMapPuzzle(t *testing.T) {
	log.Println(len(map[interface{}]int{
		new(int):      1,
		new(int):      2,
		new(struct{}): 3,
		new(struct{}): 4,
	})) // 3
}

var (
	s1, s2 = new(struct{}), new(struct{})
)

func TestAddressOfEmptyStruct(t *testing.T) {
	log.Printf("%t %p %p\n", s1 == s2, s1, s2)

	s3, s4 := new(struct{}), new(struct{})
	log.Printf("%t\n", s3 == s4)

	s5, s6 := new(struct{}), new(struct{})
	log.Printf("%t %p %p\n", s5 == s6, s5, s6)
}
