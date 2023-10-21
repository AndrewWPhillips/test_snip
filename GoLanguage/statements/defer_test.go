package __

import (
	"fmt"
	"log"
	"testing"
)

func f(i int) {
	fmt.Println(i)
}

func TestDefer1(t *testing.T) {
	a := 1
	defer f(a)
	a = 2
}

func TestDefer2(t *testing.T) {
	a := 1
	defer func() {
		log.Println(a) // prints 2
	}()
	a = 2
}

func g() func() {
	fmt.Println("In g")
	return func() {
		fmt.Println("In deferred func")
	}
}

func TestReturnDeferredFunc(t *testing.T) {
	defer g()()
	fmt.Println("In main")
}
