package __

import (
	"testing"
)

func TestEmbed(t *testing.T) {
	type Slice[T any] []T

	type X struct {
		Slice[int]
	}

	var x X
	x.Slice = nil
	_ = x
}
