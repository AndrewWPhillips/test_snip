package __

import (
	"fmt"
	"sort"
	"testing"
)

func TestSS(t *testing.T) {
	a := []string{"z", "a", "b", "abacus", ""}
	sort.Strings(a)
	fmt.Printf("%q\n", a)
}
