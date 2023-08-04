package __

import (
	"testing"
)

// requires Go 1.21

func TestMinBuiltin(t *testing.T) {
	v := []int{1, 2, 3}
	println(min(99, v...))
}
