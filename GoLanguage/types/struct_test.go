package __

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestStructUnnamedFields(t *testing.T) {
	type A struct {
		_, a, _ int
	}
	var a = A{1, 2, 3} // blank field initializers are ignored
	fmt.Println(a)     // {0 2 0}
}

func TestStructZeroSizeField(t *testing.T) {
	type A struct {
		_ [0]int
	}
	var a A
	fmt.Println(unsafe.Sizeof(a), a) // 0  {[]}
	type B struct {
		b bool
		_ [0]int
	}
	var b B
	fmt.Println(unsafe.Sizeof(b), b) // 16 {false []}
}

// TestAnonStruct2 shows that unnamed struct types with same fields are the same type
func TestAnonStruct2(t *testing.T) {
	fff := func(a map[struct {
		int
		string
	}]float64,
	) {
		fmt.Println(a)
	}

	m := make(map[struct {
		int
		string
	}]float64)

	m[struct {
		int
		string
	}{1, "a"}] = 3.14

	fmt.Println(m)
	fff(m) // map[{1 a}:3.14]

}
