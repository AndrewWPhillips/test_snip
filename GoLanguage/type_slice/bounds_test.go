package __

import (
	"fmt"
	"testing"
)

// What does this print?
// a. []
// b. [1]
// c. [1 2 3]
// d. panics
func TestEndBeforeStart(t *testing.T) {
	s := []int{0, 1, 2, 3}
	s = s[:0]
	s = s[1:]
	fmt.Println(s)
}

// d

// What does this print?
// a. [1]
// b. [1 2]
// c. [1 2 3]
// d. panics
func TestCapBeforeEnd(t *testing.T) {
	s := []int{0, 1, 2, 3}
	s = s[0:0:1]
	s = s[1:2]
	fmt.Println(s)
}

// d

// What does this print?
// a. []
// b. [1 2 3]
// c. [1 2 3 4]
// d. panics
func TestEndAfterEnd(t *testing.T) {
	s := []int{0, 1, 2, 3}
	s = s[:0]
	s = s[1:4]
	fmt.Println(s)
}

// b
