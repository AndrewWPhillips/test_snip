package __

import (
	"log"
	"testing"
)

func TestShortDecl(t *testing.T) {
	a := 1
	//a : 2 // error: no new declarations
	a, b := 2, "two" // OK but only b is a new declaration
	log.Println(a, b)
}
