//go:build go1.22

package __

import (
	"testing"
)

func TestForRangeInt(t *testing.T) {
	for i := range 10 {
		println(i)
	}
}
