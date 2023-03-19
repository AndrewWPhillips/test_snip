package __

import (
	"log"
	"testing"
)

func TestPassByValueAndPointer(t *testing.T) {
	type a struct{ b int }
	d := func(e a) { log.Println(e.b) }
	g := func(h *a) { log.Println(h.b) }
	d(a{b: 1})
	g(&a{b: 2})
}
