package __

import (
	"fmt"
	"log"
	"testing"
)

func TestBinary(t *testing.T) {
	log.Printf("%#x\n", 0b_1010_1011_1100)
}

// TestConstFloat demonstrates how compile-time constants are evaluated exactly.  That is,  f2*f2 is
// evaluated at run-time, f1*f1 uses IEEE float64 (base-2 exponent) so cannot represent 1.3 exactly
func TestConstFloat(t *testing.T) {
	const f1 = 1.3
	var f2 = 1.3
	fmt.Println(f1*f1 == f2*f2)
}
