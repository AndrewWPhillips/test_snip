package __

import (
	"fmt"
	"strings"
	"testing"
	"unique"
	"unsafe"
)

func TestUniqueString(t *testing.T) {
	u1 := unique.Make[string]("abc")
	u2 := unique.Make("abc")
	u3 := unique.Make("xyz")

	println(u1 == u2)
	println(u1 == u3)
	fmt.Printf("%v %T\n", u1.Value(), u1.Value())
}

func TestString(t *testing.T) {
	var h [10]unique.Handle[string]
	{
		// string creation from GoLanguage/type_string/internal_test.go:10
		sb := strings.Builder{}
		sb.WriteString("abc")
		ab := "ab"
		d := ab + "c"

		h[0] = unique.Make("abc")
		h[1] = unique.Make(sb.String())
		h[2] = unique.Make("abcd"[:3])
		h[3] = unique.Make(d)
		h[4] = unique.Make("xyz")
	}
	fmt.Printf("%v %v %v %v %v\n", h[0], h[1], h[2], h[3], h[4])
	fmt.Printf("%t %t\n", h[0] == h[1], h[0] == h[4])
}

// TestSubString checks if a string can be interned as part of another string
func TestSubString(t *testing.T) {
	a := unique.Make("abc")
	b := unique.Make("ab")
	c := unique.Make("abc"[:2])
	d := unique.Make("abc"[:])

	fmt.Printf("%p %p\n", unsafe.StringData("abc"), unsafe.StringData("ab"))
	fmt.Printf("%p %p\n", unsafe.StringData("abc"), unsafe.StringData("abc"[1:2]))

	// Go 1.23 there are 2 different strings interned as separate strings ("abc" and "ab")
	fmt.Printf("%p %p\n", unsafe.StringData(a.Value()), unsafe.StringData(b.Value()))
	fmt.Printf("%p %p\n", unsafe.StringData(c.Value()), unsafe.StringData(d.Value()))
}
