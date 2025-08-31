package __

import (
	"math"
	"testing"
)

// TestOptCompare check if comparing nan to itself gets optimised away
func TestOptCompare(t *testing.T) {
	n := math.NaN()
	if n != n {
		println("NaN")
	}
}
