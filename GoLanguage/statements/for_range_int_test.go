//go:build go1.22

// range over int new in Go 1.22

package __

import (
	"testing"
)

func TestForRangeInt0(t *testing.T) {
	for range 10 {
		println('x')
	}
}

func TestForRangeInt(t *testing.T) {
	for i := range 10 {
		println(i)
	}
}
