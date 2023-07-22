package __

import (
	"fmt"
	"testing"
)

type (
	Sizer interface {
		[1]int | [2]int | [3]int
	}
)

func F[I Sizer]() {
	var dummy I
	fmt.Println("Generic parameter is ", len(dummy))
	//const problem = len(dummy)
}

func TestSizer(t *testing.T) {
	F[[2]int]() // pass 2 as generic parameter
}
