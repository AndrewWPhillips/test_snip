package __

import (
	"log"
	"testing"
)

// TestNewInt looks at the address of an int on the heap
func TestNewInt(t *testing.T) {
	i := 0
	log.Printf("%p\n", &i)
	var pi *int
	log.Printf("%p %p\n", pi, &pi)
	pi = new(int)
	log.Printf("%p %d\n", pi, *pi)
	*pi = 3
	log.Printf("%p %d\n", pi, *pi)
}
