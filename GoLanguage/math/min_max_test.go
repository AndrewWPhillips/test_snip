//go:build go1.21

package __

import (
	"testing"
)

// min requires Go 1.21

func TestMinBuiltin(t *testing.T) {
	v := []int{1, 2, 3}
	println(min(99, v[0], v[1], v[2]))
}
