package __

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// TestNulByteInString tests using a nul (0) byte as a separator in a string
// I needed a map key composed of a variable number of strings so creating a single string
// using a nul byte separator worked well (assuming the strings don't contain a nul byte)
func TestNulByteInString(t *testing.T) {
	orig := []string{"AA", "BB"}
	sep := string('\x00')
	got := strings.Split(strings.Join(orig, sep), sep)
	fmt.Println(len(got), got)
	if !reflect.DeepEqual(got, orig) {
		t.FailNow()
	}
}
