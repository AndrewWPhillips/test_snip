package __test

import (
	"fmt"
	"testing"
)

func TestStackIncrease(t *testing.T) {
	a := 42

	fmt.Printf("a: %p\n", &a)
	b(&a)
	fmt.Printf("a: %p\n", &a)
}

func b(p *int) {
	var big [20]int

	fmt.Printf("%p\n", &big[0])
	fmt.Printf("%p\n", &big[1])
	fmt.Printf("p: %p\n", p)
}
