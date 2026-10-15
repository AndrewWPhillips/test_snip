package __

import (
	"fmt"
	"testing"
)

func TestLoopDefer(t *testing.T) {
	for i := range 3 {
		defer fmt.Println(i)
	}
}

func TestLoopDeferPointer(t *testing.T) {
	for i := range 3 {
		defer fmt.Println(&i) // different addresses after Go 1.22
	}
}
