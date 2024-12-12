//go:build go1.20

package __

// Uses generic slices package - new in Go 1.20

import (
	"slices"
	"testing"
)

func TestDupeString(t *testing.T) {
	in := []string{"xxx", "yyy", "a", "zzz", "b", "a", "yyy"}
	for i, g := range in {
		i2 := slices.Index(in[i+1:], g)
		if i2 != -1 {
			println("Dupe:", g, i, i+1+i2)
		}
	}
}
