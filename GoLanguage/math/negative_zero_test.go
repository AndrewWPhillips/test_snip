package __

import (
	"testing"
)

func TestNegZeroLiteral(t *testing.T) {
	x := -0.0       // Unlike C++ etc this does not work - gives +0
	println(x)      // +0
	println(x * -1) // -0
}

func TestNegZeroCalc(t *testing.T) {
	x := 0.0
	y := x * -1.0          // THIS WORKS!!!
	println(x, y)          // 0, -0
	println(y < x, y == x) // false, true
}
