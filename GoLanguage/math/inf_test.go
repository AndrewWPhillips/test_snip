package __

import (
	"math"
	"testing"
)

func TestInf(t *testing.T) {
	a := math.MaxFloat64
	println(a + a)

	b := math.SmallestNonzeroFloat64
	println(b, b-b)
}
