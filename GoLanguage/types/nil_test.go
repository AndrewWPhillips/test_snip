package __

import (
	"log"
	"testing"
)

func TestNilCompare(t *testing.T) {
	var f, g *int
	var h interface{}
	var i interface{} = f

	log.Println(g == f)          // true
	log.Println(h == f)          // false - h is not of *int type
	log.Println(i == f)          // true - both of same type (*int) with nil value
	log.Println(g == nil)        // true
	log.Println(h == nil)        // true
	log.Println(i == nil)        // false - i itself is nopt nil but contains *int type with nil value
	log.Println(i.(*int) == nil) // true
}
