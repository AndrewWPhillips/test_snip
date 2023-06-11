package __

import (
	"reflect"
	"testing"
)

func TestFuncEquality(t *testing.T) {
	a := TestFuncEquality
	b := TestFuncEquality
	t.Log(reflect.DeepEqual(a, b)) // DeepEqual considers a func not to be equal to itself
	a = nil
	b = nil
	t.Log(reflect.DeepEqual(a, b)) // DeepEqual only returns true for functions if they are both nil!!!
}

func p(a, b interface{}) {
	fmt.Print(":", reflect.DeepEqual(a, b))
}

func TestFuncArraySlice(t *testing.T) {
	a := [1]func(){func() {}} // array
	p(a, a)                   // false: arrays of func give false even if they contain the same func (see TestFuncEquality above)
	p(a[:], a[:])             // true: comparing slices with the same underlying array (+len/cap)
	b := a                    // copy the array
	p(a[:], b[:])             // false: slices with different underlying arrays
}
