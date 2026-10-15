package __

import (
	"fmt"
	"testing"
)

func TestBuiltinNew(t *testing.T) {
	type s struct{ b bool }
	var i int

	pi1 := new(int)
	pin := make([]*int, 0, 30000)
	for n := range cap(pin) {
		// The first 845 seem to be on the stack
		pin = append(pin, new(int))
		if n%100 == 0 {
			fmt.Printf("%d %p\n", n, pin[n])
		}
	}
	fmt.Printf("%p %p\n", &i, pi1)
	fmt.Printf("%p, %p, %p\n", pin[0], pin[1], pin[29999])

	ps1 := new(s)
	ps2 := &s{}

	fmt.Printf("%p %p\n", ps1, ps2)
}
