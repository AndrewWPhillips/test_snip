package __

import (
	"log"
	"testing"
	"unsafe"
)

type elt int8

func (p *elt) inc() *elt {
	*p++
	return p
}

func (p *elt) dec() *elt {
	*p--
	return p
}

// TestModifyMapElt modifies map elements using pointer receiver
func TestModifyMapElt(t *testing.T) {
	var p *elt
	log.Println(unsafe.Sizeof(p))

	m := make(map[int]elt)
	m[1] = elt(1)
	m[2] = elt(42)

	log.Println(m)

	for k, v := range m {
		// m[k].inc()  // error: cannot call pointer method on map element
		v.inc().inc().dec()
		m[k] = v
	}

	log.Println(m) // map[1:2 2:43]
}

func TestModifyArrayElt(t *testing.T) {
	a := []elt{1, 42}
	log.Println(a)

	for i, v := range a {
		_, _ = i, v
		//v.inc()    // a == [1 42]
		a[i].inc() // a == [2 43]
	}
	log.Println(a)
}
