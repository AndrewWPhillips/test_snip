package __

import (
	"fmt"
	"log"
	"testing"
)

func deferredFunction(n int) {
	log.Println("in deferredFunction, n =", n)
}

func deferredReturner(m int) func(int) {
	log.Println("in deferredReturner, m =", m)
	return deferredFunction
}

func TestDeferFunc(t *testing.T) {
	i := 1
	defer deferredFunction(i)
	i++
	log.Println("END TestDeferFunc")
}

func TestDeferLiteral(t *testing.T) {
	i := 1
	defer func(n int) {
		log.Println("in deferredFunction, n =", n)
	}(i)
	i++
	log.Println("END TestDeferLiteral")
}

func TestDeferCapture(t *testing.T) {
	i := 1
	defer func() {
		log.Println("in deferredFunction, i =", i)
	}()
	i++
	log.Println("END TestDeferCapture")
}

func TestDeferReturner(t *testing.T) {
	i := 1
	defer deferredReturner(i)(i)
	i++
	log.Println("END TestDeferReturner")
}

func TestDeferDirectRecoverCall(t *testing.T) {
	defer recover() // it seems recover is called here before panic so panic is not "caught" and test fails
	fmt.Println("next: panic 1")
	panic("test 1")
	log.Println("END TestDeferDirectRecoverCall (not printed due to above panic)")
}

func TestDeferDirectPrintlnCall(t *testing.T) {
	defer println("in println")
	fmt.Println("next: panic 2")
	panic("test 2")
	log.Println("END TestDeferDirectPrintlnCall (not printed due to above panic)")
}

func TestDeferFmtPrintlnCall(t *testing.T) {
	defer fmt.Println("in fmt.Println")
	fmt.Println("next: panic 3")
	panic("test 3")
	log.Println("END TestDeferFmtPrintlnCall (not printed due to above panic)")
}

func recoverFunction() {
	log.Println("recover =", recover())
}

func TestDeferRecoverFunc(t *testing.T) {
	defer recoverFunction() // it seems recover is called here (not after return) but it is hard to tell
	fmt.Println("next: panic 1A")
	panic("test 1A")
	log.Println("END TestDeferRecoverFunc (not printed due to above panic)")
}

func TestDeferRecoverLiteral(t *testing.T) {
	defer func() {
		recover()
	}()
	fmt.Println("next: panic 1B")
	panic("test 1B")
	log.Println("END TestDeferRecoverLiteral (not printed due to above panic)")
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

func getDeferredFunc(i int) func(int) {
	fmt.Println("in getDeferredFunc:", i)
	return func(i int) {
		fmt.Println("in deferred func:", i)
	}
}

func TestEvalOrder2(t *testing.T) {
	i := 1
	defer getDeferredFunc(i)(i)
	i = 2
	fmt.Println("in main:", i)
}
