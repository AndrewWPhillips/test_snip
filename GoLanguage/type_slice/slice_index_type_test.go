package __

import (
	"testing"
)

// TestIndexTypes tests what int types can be used to index into a slice
func TestIndexTypes(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e", "f"}

	type MyInt int64
	type MyInt2 struct{ int }

	var i8 int8 = 1
	var u uint = 2
	var p uintptr = 3
	var m MyInt = 4
	var m2 MyInt2 = MyInt2{int: 5}

	t.Log(s[i8])
	t.Log(s[u])
	t.Log(s[p])
	t.Log(s[m])
	//t.Log(s[m2])
	t.Log(s[m2.int])
}
