package __

import (
	"fmt"
	"testing"
)

// TestPanic causes a nil reference
func TestPanic(t *testing.T) {
	var ss *struct{ i int }
	fmt.Println(ss.i)
}
