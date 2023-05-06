package __

import (
	"log"
	"testing"
)

func TestInterfaceAssign(t *testing.T) {
	a := 1
	i1 := interface{}(a)
	a = 2
	log.Println(i1)
	i2 := i1
	i1 = a
	log.Println(i2)
	log.Println(i1)
}

func TestInterfaceReassign(t *testing.T) {
	var i1 interface{}
	i1 = 1
	i2 := i1
	i2 = 2 // since i1 and i2 are references
	log.Println(i1, i2)
}
