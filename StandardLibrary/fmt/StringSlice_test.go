package __

import (
	"fmt"
	"testing"
)

func TestStringSlice(t *testing.T) {
	a := []string{"string 1", "string 2"}
	b := struct{ string }{"string_3"}
	fmt.Printf("%s\n", a) // [string 1 string 2]
	fmt.Printf("%q\n", a) // ["string 1" "string 2"]
	fmt.Printf("%v\n", a) // [string 1 string 2]
	fmt.Printf("%s\n", b) // {string_3}
	fmt.Printf("%q\n", b) // {"string_3"}
	fmt.Printf("%v\n", b) // {string_3}
}
