package __

import (
	"fmt"
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

// TestRuneCast tests length of rune slice (cast from string)
func TestRuneCast(t *testing.T) {
	const s = "touché"
	r := []rune(s)
	b := []byte(s)
	log.Println(len(s), len(r), len(b))
}

func TestLen(t *testing.T) {
	const s = "123456789"
	var a byte = 1 << len(s) / 128
	var b byte = 1 << len(s[:]) / 128
	// len(s[:]) is evaluated at runtime while len(s) is a compile-time constant
	// the first expression is evaluated at compile time (using numeric arithmetic) whereas
	// the 2nd expression is evaluated at run-time using byte arithmetic - it does NOT
	// use int arithmetic since the only non-constant is on the right side of a shift operation
	fmt.Println(a, b)
}
