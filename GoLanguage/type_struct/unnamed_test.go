package __

import (
	"fmt"
	"testing"
)

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
