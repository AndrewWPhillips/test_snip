//go:build go1.22

package __

import (
	"testing"
)

func TestForRangeInt(t *testing.T) {
	for i := range 100 {
		println(i)
	}
}
