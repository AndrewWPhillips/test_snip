package __

import (
	"log"
	"strings"
	"testing"
	"unsafe"
)

// TestMutate demonstrates how to mutate a string.
// NOT RECOMMENDED: This is unsafe code (in many ways) since a lot of things assume strings are immutable.
func TestMutate(t *testing.T) {
	//s := "abc"  // This does not work as string literals are in read-only memory

	// s := string([]byte{'a', 'b', 'c'})

	// build a string (to ensure it's not in read-only memory)
	sb := strings.Builder{}
	sb.WriteString("abc")
	s := sb.String()

	b := *(*struct {
		p      *[4000000000]byte
		length int
	})(unsafe.Pointer(&s))
	// Be very careful not to write past the end of the string
	b.p[b.length-1]++ // inc last char (c -> d)
	log.Printf("%p %p %v %q", &s, &b, b.p[0], s)
}
