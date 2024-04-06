package __

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestFuncSize(t *testing.T) {
	f := func() {}
	fmt.Println(unsafe.Sizeof(f))

	i := 42
	g := func(j int) int { return i + j }
	fmt.Println(unsafe.Sizeof(g))
}
