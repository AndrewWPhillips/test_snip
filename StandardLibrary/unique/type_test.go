package __

import (
	"fmt"
	"testing"
	"unique"
)

func TestSameUnderlyingTypes(t *testing.T) {
	type Str string

	h1 := unique.Make("abc")
	h2 := unique.Make(Str("abc"))
	fmt.Println(h1, h2) // different pointers
}

func TestUniqueInterface(t *testing.T) {
	var a any

	a = []int{}
	unique.Make(a) // panic as slices aren't comparable

	_ = map[any]struct{}{a: struct{}{}} // panic for the same reason
}

func TestUniqueInterface2(t *testing.T) {
	var a, b interface{}

	a, b = 1, 2
	h1 := unique.Make(a)
	h2 := unique.Make(b)
	fmt.Println(h1, h2)
}
