package __

import (
	"fmt"
	"testing"
)

type (
	A struct {
		B int
	}
	C struct {
		A
		D string
	}
)

func TestInitEmbedded(t *testing.T) {
	x := C{
		A: A{B: 1},
		D: "test",
	}
	fmt.Println(x)
}

type (
	E1 struct {
		i int
	}
	E2 struct {
		i, e, f int
	}
	Outer struct {
		E1
		E2
		e, g int
	}
)

func TestEmbedConflict(t *testing.T) {
	v := Outer{
		E1: E1{
			i: 11,
		},
		E2: E2{
			i: 22,
			e: 33,
			f: 44,
		},
		e: 55,
		g: 66,
	}

	//println(v.i) // ERROR: ambiguous selector v.i
	println(v.E1.i) // 11
	println(v.E2.i) // 22
	println(v.e)    // 55 [Outer]
	println(v.E2.e) // 33 [E2]
	println(v.f)    // 44
	println(v.g)    // 66
}
