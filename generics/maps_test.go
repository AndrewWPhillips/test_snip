package __

import (
	"log"
	"math"
	"testing"

	"golang.org/x/exp/maps"
)

func TestMapsEqualNan(t *testing.T) {
	a := map[float64]struct{}{math.NaN(): struct{}{}}
	b := map[float64]struct{}{math.NaN(): struct{}{}}
	log.Println(maps.Equal(a, b)) // false!
}

type pointerStruct struct{ p *int }

// TestMapsEqualDepth tests how "deep" the equality test of maps.Equal is
// As expected it only compares each element using a shallow comparison
func TestMapsEqualDepth(t *testing.T) {
	i, j := 42, 42
	m, n := make(map[int]pointerStruct), make(map[int]pointerStruct)

	m[0] = pointerStruct{p: &i}
	n[0] = pointerStruct{p: &j}
	log.Println(maps.Equal(m, n)) // false - the p fields have different addresses (even though they both point to 42)
}
