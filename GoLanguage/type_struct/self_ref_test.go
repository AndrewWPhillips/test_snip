package __

import (
	"log"
	"testing"
)

func TestSelfRefCopy(t *testing.T) {
	var a, b struct {
		i int
		p *int
	}

	a.i = 2
	a.p = &a.i

	b.i = 3
	b.p = &b.i

	b = a
	log.Println(*b.p)
	b.i = 4
	log.Println(*b.p)
}

func New() (r struct {
	i int
	p *int
}) {
	r.i, r.p = 1, &r.i
	return r
}

func TestSelfRefCopy2(t *testing.T) {
	v := New()
	v.i = 2
	log.Println(*v.p)
}
