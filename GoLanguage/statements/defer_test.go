package __

import (
	"log"
	"testing"
)

// TestDefer tests something we all should know
func TestDefer(t *testing.T) {
	a := 1
	defer func() {
		log.Println(a) // prints 2
	}()
	a = 2
}

// TestDeferReturn tests changing the return value inside a defer statement
func TestDeferReturn(t *testing.T) {
	f := func() (ret int) {
		defer func() { ret = 42 }()
		return 1
	}
	log.Println(f())
}

// TestDeferReturnWrong shows the wrong way defer statement
func TestDeferReturnWrong(t *testing.T) {
	f := func() int {
		ret := 1
		defer func() {
			ret = 42 // by the time ret is set here it's value has already been used as the return value
		}()
		return ret
	}
	log.Println(f())
}

// TestEvalOrder tests when things are executed if defer func returns a func
func TestEvalOrder(t *testing.T) {
	f := func(f func()) func(f func()) {
		f()
		return func(f func()) {
			println("A")
		}
	}

	defer f(func() { print("B") })(func() { print("D") })

	print("C")
}
