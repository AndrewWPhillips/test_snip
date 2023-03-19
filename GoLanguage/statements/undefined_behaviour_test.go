package __

import (
	"testing"
)

// TestUndef tries (and fails) to use side-effects in expressions (as in C) to get undef behaviour
func TestUndef(t *testing.T) {
	a := [...]int{11, 12, 13, 14}
	i := 1
	j := 2
	_ = j
	//a[i++] = i    // illegal
	//a[i] = i++    // illegal
	//a[i+=1] = i   // illegal
	//j = i = 1     // illegal
	i++         // statement not expression
	i += 1      //    "             "
	a[i] = i    // no way to have undef. behaviour due to side-effects in expressions
	print(a[3]) // 3
}
