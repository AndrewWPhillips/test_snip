package __

import (
	"log"
	"math/big"
	"testing"
)

// TestBigCompare shows how to compare big ints
func TestBigCompare(t *testing.T) {
	a := big.NewInt(1)
	b := big.NewInt(2)
	c := big.NewInt(2)

	log.Println(a == b, b == c) // wrong since big.Int is a ptr
	//log.Println(*a == *b, *b == *c) // compare structs containing a slice won't work
	log.Println(a.Cmp(b) == 0, b.Cmp(c) == 0) // OK

}
