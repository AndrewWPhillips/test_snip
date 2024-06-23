package __

// Demonstrates how we would use a generic interface if we want a
// generic Less method.  (Why not use a generic function?)

import (
	"testing"
)

/*
type Lesser interface {
	Less(rhs Lesser) bool
}

func IsLess[T Lesser](lhs, rhs T) bool {
	return lhs.Less(rhs)
}
*/

type Lesser[T any] interface {
	Less(T) bool
}

func IsLess[T Lesser[T]](x, y T) bool {
	return x.Less(y)
}

type Int int

// Less allows Int to be a Lesser
func (lhs Int) Less(rhs Int) bool {
	return lhs < rhs
}

func TestIsLess(t *testing.T) {
	println(IsLess(Int(10), 20))
}
