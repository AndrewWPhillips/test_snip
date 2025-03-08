package __

import (
	"fmt"
	"testing"
	"unique"
)

func TestArray(t *testing.T) {
	a := [2]int{0, 1}
	ha := unique.Make(a)
	hb := unique.Make([2]int{0, 1})
	hc := unique.Make([2]int{})
	a[1] = 0
	hd := unique.Make(a)
	fmt.Println(ha == hb, ha == hc, ha == hd) // T F F
	fmt.Println(hb == hc, hb == hd)
	fmt.Println(hc == hd) // T
}
