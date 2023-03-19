package __

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

// TestCompare tests different ways to compare strings
// see also GoLanguage/types/string_test.go/TestBinCompare
func TestCompare(t *testing.T) {
	const a = "Andrew"
	b := strings.ToLower(a)

	log.Println(a == b)                  // false
	log.Println(reflect.DeepEqual(a, b)) // false
	log.Println(strings.Compare(a, b))   // -1
	log.Println(strings.Compare(a, a))   // 0
	log.Println(strings.Compare(b, a))   // 1
	log.Println(strings.EqualFold(a, b)) // true
}
