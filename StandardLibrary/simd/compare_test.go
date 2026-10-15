package __test

import (
	"testing"
)

func TestInt8Equality(t *testing.T) {
	x, y := RandVector8x32(4), RandVector8x32(4)

	println(x.String())
	println(y.String())
	println(x.Equal(x).ToInt8x32().String())
	println(x.Equal(y).ToInt8x32().String())
}
