package __

import (
	"fmt"
	"log"
	"math"
	"testing"
)

// TestNanCompare shows that comparing nans is not "reflexive"
func TestNanCompare(t *testing.T) {
	nan := math.NaN()
	nan2 := nan
	fmt.Println(nan == nan2)
	fmt.Println(nan != nan2)
}

// See also TestNanKey in GoLanguage/maps/key_test.go

func TestZeroTimesInf(t *testing.T) {
	x := 0.
	x *= 1. / x
	log.Println(x)
}

// a. 0
// b. +Inf
// c. NaN
// d. panic: runtime error: divide by zero
// e. does not compile
