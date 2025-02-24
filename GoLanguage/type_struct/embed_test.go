package __

import (
	"fmt"
	"testing"
)

type (
	A struct {
		B int
	}
	C struct {
		A
		D string
	}
)

func TestInitEmbedded(t *testing.T) {
	x := C{
		A: A{B: 1},
		D: "test",
	}
	fmt.Println(x)
}
