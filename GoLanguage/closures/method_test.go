package __

import (
	"fmt"
	"log"
	"testing"
)

type myInt int

func (i myInt) print() { fmt.Println(i) }

type myFloat float64

func (f myFloat) print() { fmt.Println(f) }

func TestReceiverCapture(t *testing.T) {
	// a closure is not just a function pointer but can also have data (like a C# delegate)
	a := myInt(3)
	b := a + 4
	f := a.print
	f() // 3
	g := f
	f = b.print
	f() // 7
	g() // 3
	f = func() { fmt.Println(42) }
	f() // 42

	f = myFloat(99).print
	f() // 99
}

type pair struct {
	a int
	b int
}

func (p pair) m(i int) {
	log.Println("pair", p.a, p.b, i)
}

func f(i int) {
	log.Println("f", i)
}

func TestSameMethodAndFunc(T *testing.T) {
	c := pair{1, 2}.m // assign method to closure
	c(3)
	c = f // assigning a func to same var it must have the same signature: func(int)
	c(4)
}

func TestPtr(t *testing.T) {
	p := &pair{1, 2}
	p.m(3) // 1 2 3
	p = nil
	//p.m(5) // nil ptr deref

	pair{1, 2}.m(6)    // 1 2 6
	(&pair{1, 2}).m(7) // 1 2 7
}
