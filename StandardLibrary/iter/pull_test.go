package __

import (
	"iter"
	"testing"
)

func TestPullOne(t *testing.T) {
	next, stop := iter.Pull(OneRange(5))
	defer stop()

	println(next())
	println(next())
	println(next())
	println(next())
	println(next())
	println(next())
}
