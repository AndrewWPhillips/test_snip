package __

import (
	"fmt"
	"math"
	"testing"
)

func TestNanCompare(t *testing.T) {
	nan := math.NaN()
	nan2 := nan
	fmt.Println(nan == nan2)
	fmt.Println(nan != nan2)
}

// See also TestNanKey in GoLanguage/maps/key_test.go
