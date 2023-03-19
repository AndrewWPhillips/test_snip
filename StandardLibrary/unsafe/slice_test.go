package __

import (
	"log"
	"testing"
	"unsafe"
)

// TestModifyString shows that strings are not immutable if you use unsafe stuff
// see also StandardLibrary/string/mutate_test.go
func TestModifyString(t *testing.T) {
	s := "Hello, world"
	b := []byte(s) // Copies all bytes of the string to a new byte slice
	b[0] = 'h'

	s = *(*string)(unsafe.Pointer(&b)) // create string from slice w/o copy
	log.Println(s)
	b[0] = 'j' // WARNING: the string is now mutable via b
	log.Println(s)
}
