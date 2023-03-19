package __

import (
	"log"
	"testing"
)

type aa struct{}

func (this aa) Cc() { return }

type bb struct{}

func (this bb) Cc(i int) int { return i }

type ii interface {
	Cc() int
}

// TestInterfaceMatch tests how method signature matching works
// TODO: this was some early code I wrote - could be simpler and more comprehensive
func TestInterfaceMatch(t *testing.T) {
	a := new(aa)
	log.Println(a)
	b := new(bb)
	log.Println(b)

	var i ii
	// i = a  // build error: have Cc() want Cc() int
	// i = b  // build error: have Cc(int) int want Cc() int
	log.Println(i)

	var j interface{}
	j = a
	// i = j  // build error: j does not implement Cc() int
	log.Println(j)
}
