package __

import (
	"log"
	"testing"
)

type Restriction1 interface {
	ImplementsRestriction1()
}

type Imp1A int

func (Imp1A) ImplementsRestriction1() {} // this just says that Imp1A type satisfies Restriction1

// restrict can only be passed a value that implements Restriction1
func restrict(c Restriction1) {
	switch v := c.(type) {
	case Imp1A:
		log.Printf("ImplA: %d\n", int(v))
	}
}

// TestRestrictInterface tests using an interface to limit the acceptable parameters to a function
func TestRestrictInterface(t *testing.T) {
	restrict(Imp1A(2))

	// restrict(3)  // cannot use 3 (type int) as type Restriction1 in argument to restrict
}
