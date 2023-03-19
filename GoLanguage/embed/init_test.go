package __

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
