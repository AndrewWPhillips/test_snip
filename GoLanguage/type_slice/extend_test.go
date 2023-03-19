package __

import (
	"log"
	"testing"
)

// TestExtendFront shows that capacity is asymmetrical - can't reslice with -ve start
func TestExtendFront(t *testing.T) {
	a := [...]int{1, 2, 3}
	s := a[1:3]
	log.Printf("%v\n", s) // 2, 3
	// slice only stores capacity - doesn't know there are available bytes before start
	//s = s[-1:]                  // Compile error: invalid slice index -1 (index must be non-negative)
	//log.Printf("%v\n", s)       // 1, 2, 3
}

func TestExtend(t *testing.T) {
	a := [...]int{1, 2, 3}
	s := a[:2]
	log.Printf("%v\n", s) // 1, 2
	s = s[:len(s)+1]      // can extend a slice by reslice up to cap
	log.Printf("%v\n", s) // 1, 2, 3
	//s = s[:len(s)+1]              // run-time error: slice bounds
	s = append(s, 0)
	log.Printf("%v\n", s) // 1, 2, 3, 0
}
