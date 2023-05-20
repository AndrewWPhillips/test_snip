package __

import (
	"log"
	"math/big"
	"testing"
)

func TestCompare(t *testing.T) {
	var a big.Int
	a.SetString("42", 10)
	log.Println(a.Cmp(big.NewInt(41))) // 1
	log.Println(a.Cmp(big.NewInt(42))) // 0
	log.Println(a.Cmp(big.NewInt(43))) // -1
}
