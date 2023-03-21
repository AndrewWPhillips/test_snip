package __

import (
	"log"
	"testing"
)

type Aunmix struct{}

type Bunmix struct {
	Aunmix
	c int
}

func unmix(a *Aunmix) {
	log.Printf("%T %v\n", a, a)
}

func TestUnmix(t *testing.T) {
	// test to see if we can restrict types passed to a function by "deriving" using mixin.
	// Ie See if we can pass a ptr to Bunmix when a ptr to the base class (Aunmix) is expected
	b := Bunmix{Aunmix: Aunmix{}, c: 2}
	//unmix(&b)       // no good: ptr to b to not convert to ptr to "base" type
	unmix(&b.Aunmix) // OK

}
