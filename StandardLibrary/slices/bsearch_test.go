//go:build go1.22

package __

import (
	"math"
	"slices"
	"testing"
)

func TestBig(t *testing.T) {
	x, ok := slices.BinarySearchFunc( // added in Go 1.22
		make([]struct{}, math.MaxInt),
		struct{}{},
		func(elem, target struct{}) int { return -1 },
	)
	if x != math.MaxInt || ok {
		t.Errorf("TestBig: got %v %v, want %v %v", x, ok, math.MaxInt, false)
	}
}
