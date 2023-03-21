package __

import (
	"log"
	"testing"
)

func TestMapModifyDuringIteration(t *testing.T) {
	//	m := map[int]struct{}{0: {}, 1e9: {}, 2e9: {}}
	m := make(map[int]struct{}, 30)
	m[0] = struct{}{}
	m[1e9] = struct{}{}
	m[2e9] = struct{}{}
	for i := range m {
		n := i + 1
		log.Printf("got %d (len = %d) adding %d\n", i, len(m), n)
		m[n] = struct{}{}
	}
}

func TestModifyValue(t *testing.T) {
	type s struct{ i int }

	a := new(s)
	a.i = 42
	b := new(s)
	b.i = 73

	m := make(map[string]*s)
	m["a"] = a
	m["b"] = b

	p := m["a"]
	p.i = 1
	log.Println(m["a"])
}
