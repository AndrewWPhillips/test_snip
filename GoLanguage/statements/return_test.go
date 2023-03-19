package __

import (
	"log"
	"testing"
)

func TestMultReturn(t *testing.T) {
	// CONCLUSION: for a function returning multiple values you can ignore them all but
	//             if you want just the first one you have to assign ignored ones to _
	//             This is different to the behaviour of multiple return types for
	//             map element access, chan reads, interface type conversion (T)
	//var ff func() (int, bool)
	ff := func() (int, bool) { return 2, true }

	ii, bb := ff() // OK assign both return values
	ff()           // OK ignore return values
	//ii = ff()      // ERROR: mult value in single value context
	ii, _ = ff() // OK

	log.Println(ii, bb) // 2 true
}
