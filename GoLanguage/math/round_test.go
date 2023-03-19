package __

import (
	"fmt"
	"math"
	"testing"
)

func TestApprox(t *testing.T) {
	testData := map[string]struct {
		a, b float64
	}{
		"z0": {0, 0},
		"z1": {0, 0.000001},
		"z2": {0.000001, 0},
		"z3": {0.000001, 0.00001},
		"s0": {0, 1},
		"s1": {0, 0.1},
		"s2": {0, 0.01},
		"s3": {0, 0.001},
		"s4": {0, 0.0001},
		"s5": {0, 0.00001},
		"s6": {0, 0.000001},
		"d0": {1, 2},
		"d1": {1, 1.1},
		"d2": {1, 1.01},
		"d3": {1, 1.001},
		"d4": {1, 1.0001},
		"d5": {1, 1.00001},
		"d6": {1, 1.000001},
		"b0": {1e9, 2e9},
		"b1": {1e9, 1.1e9},
		"b2": {1e9, 1.01e9},
		"b3": {1e9, 1.001e9},
		"b4": {1e9, 1.0001e9},
		"b5": {1e9, 1.00001e9},
		"b6": {1e9, 1.000001e9},
	}

	for name, d := range testData {
		d := d
		t.Run(name, func(t *testing.T) {
			fmt.Printf("%40s: %f %f %t\n", "", d.a, d.b, Approx(d.a, d.b))
		},
		)
	}
}

// Approx checks if 2 numbers are close to 3 decimal places
// Note that this is good for "non-scientific" numbers like money
// For scientific numbers you need to take into account the scale (so that 1e30 and 1.000001e30 are approx equal)
func Approx(a, b float64) bool {
	return math.Abs(a-b) < 0.001
}
