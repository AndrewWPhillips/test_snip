package __

import (
	"log"
	"testing"
)

const (
	a = 1
	b
	c = iota
	d
)

func TestIota1(t *testing.T) {
	log.Println(a + b + c + d)
}
