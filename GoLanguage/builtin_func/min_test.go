package __test

import (
	"log"
	"testing"
)

func TestMinTypes(t *testing.T) {
	f := float64(1)
	c := min(f, 2, 3)
	log.Printf("%T %v\n", c, c)

	d := min(42, 3.14, 'a')
	log.Printf("%T %v\n", d, d)
}
