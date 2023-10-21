package __

import (
	"log"
	"testing"
)

// TestNoContract checks the sort of error messages you might get when using a specific type as a type
// parameter that does not support a specific operation need by the generic function/type.
func TestNoContract(t *testing.T) {
	a := struct {
		int
		string
	}{1, "a"}
	b := a

	//if a < b {   // invalid operation: a < b (operator < not defined on struct)
	if a == b {
		log.Println("YES")
	}

	//log.Println(float64("abc"))    // cannot convert "abc" (type untyped string) to type float64
	//log.Println(float64(a == b))   // cannot convert a == b (type bool) to type float64
}
