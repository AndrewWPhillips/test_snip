package __

import (
	"log"
	"testing"
)

// TestBinCompare tests comparison when string has nul bytes
func TestBinCompare(t *testing.T) {
	a := []byte{'a', 'b', 'c', 0, 1}
	b := []byte{'a', 'b', 'c', 0, 1}
	c := []byte{'a', 'b', 'c', 0, 255}
	d := []byte{'a', 'b', 'c', 0, 0}
	log.Println(string(a) == string(b))         // true
	log.Println(string(a) == string(c))         // false
	log.Println(string(a) < string(c))          // true
	log.Println(len(a), len(b), len(c), len(d)) // all 5
}
