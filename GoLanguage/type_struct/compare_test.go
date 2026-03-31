package __

import (
	"fmt"
	"reflect"
	"testing"
)

// TestAnonFields cannot build anymore (can't assign to _ fields)
func TestAnonFieldsCompare(t *testing.T) {
	type ts struct {
		_ int
		_ bool
	}

	var t1 = ts{1, true}
	var t2 = ts{2, false}
	fmt.Println(t1 == t2) // true
}

func TestNamedStructCompare(t *testing.T) {
	type t1 struct{ i int }
	type t2 struct{ i int }

	v1 := t1{42}
	v2 := t2{42}
	//println(v1 == v2) // Invalid operation: v1 == v2 (mismatched types t1 and t2)
	println(v1 == t1(v2))              // true
	println(reflect.DeepEqual(v1, v2)) // false!! due to diff type names
}

func TestUnnamedStructCompare(t *testing.T) {
	v1 := struct{ i int }{42}
	v2 := struct{ i int }{42}
	println(v1 == v2)                  // true
	println(reflect.DeepEqual(v1, v2)) // true
}

// TestMixedStructCompare checks if you can compare anon struct to named struct with same fields
func TestMixedStructCompare(t *testing.T) {
	type t1 struct{ i int }

	v1 := t1{42}
	v2 := struct{ i int }{42}
	println(v1 == v2)                  // true
	println(reflect.DeepEqual(v1, v2)) // false!! (one has type name)
}

func TestTaggedStructCompare(t *testing.T) {
	v1 := struct {
		i int `x`
	}{42}
	v2 := struct {
		i int `x`
	}{42}
	v3 := struct {
		i int `y`
	}{42}
	println(v1 == v2)                  // true
	println(reflect.DeepEqual(v1, v2)) // true
	//println(v1 == v3)                  // Invalid operation: v1 == v3 (mismatched types ...)
	println(reflect.DeepEqual(v1, v3)) // false due to different field metadata
}
