package __test

import (
	"fmt"
	"testing"
	"unique"
)

func TestUniqueString(t *testing.T) {
	u1 := unique.Make[string]("abc")
	u2 := unique.Make("abc")
	u3 := unique.Make("xyz")

	println(u1 == u2)
	println(u1 == u3)
	fmt.Printf("%v %T\n", u1.Value(), u1.Value())
}
