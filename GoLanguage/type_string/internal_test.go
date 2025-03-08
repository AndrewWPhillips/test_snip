package __

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

// TestStringData uses unsafe.StringData to get a pointer to the underlying bytes of a string
func TestStringData(t *testing.T) {
	a := "abc"

	sb := strings.Builder{}
	sb.WriteString("abc")
	b := sb.String()

	c := "abcd"[:3]

	ab := "ab"
	d := ab + "c"

	// 0x9...F0B, same, 0xC...4C0, 0x9...8, 0xC...4C8
	fmt.Printf("%p %p, %p, %p, %p\n", unsafe.StringData("abc"),
		unsafe.StringData(a), unsafe.StringData(b),
		unsafe.StringData(c), unsafe.StringData(d))
}
