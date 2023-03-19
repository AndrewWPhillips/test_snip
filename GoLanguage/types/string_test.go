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
	log.Println(string(a) == string(b))
	log.Println(string(a) == string(c))
	log.Println(string(a) > string(c))
	log.Println(len(a), len(b))
}
