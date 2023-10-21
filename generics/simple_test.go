package __

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints"
)

// Name is just a single value of the generic type stored in a single element array
type Name[T any] [1]T

func TestSimpleType(t *testing.T) {
	var a [1]int
	fmt.Printf("%T %v\n", a, a)
	var b Name[int]
	fmt.Printf("%T %v\n", b, b)
}

type Numeric interface {
	constraints.Integer | constraints.Float
}

// Simple is just a generic func that returns its parameter
func Simple[T Numeric](a T) T {
	return a / 2
}

func TestSimpleFunc(t *testing.T) {
	a := Simple(3.14)
	fmt.Printf("%T %v\n", a, a)
}

func Zero[T any]() (r T) {
	return
}
