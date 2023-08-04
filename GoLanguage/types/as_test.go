package __

import (
	"log"
	"testing"
)

func TestAs(t *testing.T) {
	var x interface{}
	i := 42
	x = i

	// I thought Go had as/is but that must have been C#
	// You can do the same thing with a type assertion
	as, is := x.(int)
	log.Println(as, is)
}
