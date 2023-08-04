package __

import (
	"fmt"
	"testing"
)

// TestAnonFields cannot build anymore (can't assign to _ fields)
func TestAnonFields(t *testing.T) {
	type ts struct {
		_ int
		_ bool
	}

	var t1 = ts{1, true}
	var t2 = ts{2, false}
	fmt.Println(t1 == t2)  // true
}
