package __

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	m := make(map[string]int)
	m["abc"] = 42
	log.Println(m["abc"]) // 42
	delete(m, "abc")
	log.Println(len(m)) // 0
	delete(m, "abc")
	log.Println(len(m)) // 0

	// Use of nil map
	m = nil
	v, got := m["xx"]
	log.Println(v, got) // 0 false
	//m["abc"] = 1        // panic: assignment to entry in nil map
	log.Println(len(m)) // 0
}
