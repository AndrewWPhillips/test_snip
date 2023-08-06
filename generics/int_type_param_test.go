package __

import (
	"log"
	"testing"
)

// This is a test of using arrays as type parameters to fake integer (non-type) generic parameters.
// Background: C++ supports "non-type" (integer, enum and pointer) parameters to templates (C++ templates ~= generics)
// Go does not support non-type generic parameters (which is the only thing I really miss with Go generics) so I have been
// working on a way around this by using array types as type parameters - using the size of the array as a sort of integer type parameter
// eg: type T interface( [2]int, [3]int } - can be used as constraint which when used with generic class/function can allow an array
// of size 2 or 3 to be passed as a generic (compile-time) "parameter" and the integer obtained using len(array).

// INT2 and INT3 are just used as part of type constraint T_INT
// TODO: check if T_INT actually gets instantiated much and if so maybe use [2]struct{} instead of [2]any - since
// struct{} uses 0 bytes while interface{} uses 16 (or 8 on 32-bit arch?) bytes
type INT2 [2]any
type INT3 [3]any
type T_INT interface{ INT2 | INT3 } // use array to fake integer parameter which can be either 2 or 3

func f[T T_INT]() int {
	var v T
	return len(v) // not compile-time constant (anymore) - was in Go2 (Go1.18 beta)
}

func TestArrayConstraintLen(t *testing.T) {
	log.Println(f[INT2]())
	log.Println(f[INT3]())
}
