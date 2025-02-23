package __

import (
	"fmt"
	"testing"
)

func TestForRangeArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	for _, v := range a {
		v = 13
		v++
	}
	fmt.Println(a)
}
