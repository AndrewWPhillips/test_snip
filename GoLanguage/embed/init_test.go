package __

// This embed directory is for tests of embedded structs (called mixins in some languages).
// For embedding files into the binary using //go:embed - see separate directives tests.

import (
	"log"
	"testing"
)

func TestStructInit(t *testing.T) {
	type bb struct {
		m, n int
	}
	type ss struct {
		bb
		A, B int
	}

	a := &ss{bb: bb{m: 42}, B: 3}

	log.Println(a.m) // 42
	log.Println(a.B) // 3
}
