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

const (
	sMax32      = "4294967295"
	sMax32plus1 = "4294967296"
	sMax64      = "18446744073709551615"
	sMax64plus1 = "18446744073709551616"
)

var (
	max32, max32plus1 big.Int
	max64, max64plus1 big.Int
)

func init() {
	max32.SetString(sMax32, 10)
	max32plus1.SetString(sMax32plus1, 10)
	max64.SetString(sMax64, 10)
	max64plus1.SetString(sMax64plus1, 10)
}

func TestIntAdd(t *testing.T) {
	tests := []struct {
		name   string
		p1, p2 *big.Int
		exp    *big.Int
	}{
		{"zero", big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{"1+1", big.NewInt(1), big.NewInt(1), big.NewInt(2)},
		{"1+0", big.NewInt(1), big.NewInt(0), big.NewInt(1)},
		{"0+1", big.NewInt(0), big.NewInt(1), big.NewInt(1)},
		{"oflow32", &max32, big.NewInt(1), &max32plus1},
		{"oflow64", &max64, big.NewInt(1), &max64plus1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := new(big.Int)
			got.Add(test.p1, test.p2)
			if got.Cmp(test.exp) != 0 {
				t.Fatalf("%-12s: expected %v but got %v\n", test.name, test.exp, got)
			}
		})
	}
}
