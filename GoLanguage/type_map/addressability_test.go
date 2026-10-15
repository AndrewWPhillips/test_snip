package __

import (
	"testing"
)

func TestMapElementAddressability(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2, "c": 3}
	//ptr := &a["a"] // error: cannot take the address of a["a"]
	_ = a
}

func TestSliceElementAddressability(t *testing.T) {
	//var b []int
	b := []int{1, 2, 3}
	ptr2 := &b[1] // OK to take address of slice element
	println(*ptr2)

	b = append(b, 4, 5, 6, 7)
	b[1] = 42
	println(*ptr2) // still points to a valid address (with old contents)
}

// TestStructFieldAddressability shows that you can't get the address of a field of a struct
// literal.  (A literal is not addressable hence it's fields aren't.)  BUT you can take the
// address of a slice element since a slice always points to an underlying array (or is nil).
func TestStructFieldAddressability(t *testing.T) {
	//ptr := &struct{ List []string }{List:[]string{"a"}}.List // can't take address of field
	ptr := &struct{ List []string }{List: []string{"a"}}.List[0] // OK: address of slice element

	println(*ptr) // a
}
