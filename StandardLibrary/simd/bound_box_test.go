//go:build go1.26 && goexperiment.simd && amd64

package __

import (
	"log"
	"math/rand/v2"
	"simd/archsimd"
	"testing"
)

// TestFindClosePoints demonstrates using SIMD to speed up finding points within a distance of point.
// It uses SIMD instructions to do a fast bounding box search.
func TestFindClosePoints(t *testing.T) {
	const (
		limit = 100_000 // limit of our test data coordinates
	)

	// Create lots of points within a square of size 'limit'
	// Note that we use separate arrays for the X and Y coords - TODO: check if this gives better performance
	// Also the size of the slices need to be a multiple of the AVX2 register size (32 bytes = 256 bits)
	// It only takes a few seconds to check millions of points
	var xPoints, yPoints [100_000 * 0x400]int16
	for i := range xPoints {
		xPoints[i] = int16(rand.IntN(limit))
		yPoints[i] = int16(rand.IntN(limit))
	}

	if !archsimd.X86.AVX2() {
		t.Fatal("AVX2 is required for BroadcastInt16x16(), Int16x16.Less() etc")
	}

	// Get a random point in the square
	x, y := int16(rand.IntN(limit)), int16(rand.IntN(limit))
	log.Printf("Checking: (%d, %d)\n", x, y)

	// Find all points in the square within this distance
	dist := int16(limit / 5000)
	distSq := dist * dist

	// Create SIMD vectors with X/y limits (bounding box)
	vMinX := archsimd.BroadcastInt16x16(x - dist)
	vMaxX := archsimd.BroadcastInt16x16(x + dist)
	vMinY := archsimd.BroadcastInt16x16(y - dist)
	vMaxY := archsimd.BroadcastInt16x16(y + dist)

	for i := 0; i < len(xPoints); i += 32 {
		vx := archsimd.LoadInt16x16Slice(xPoints[i : i+32])
		vy := archsimd.LoadInt16x16Slice(yPoints[i : i+32])

		// Check if point(s) are within the bounding box
		boundX := vx.Greater(vMinX).And(vx.Less(vMaxX))
		boundY := vy.Greater(vMinY).And(vy.Less(vMaxY))

		// bits := boundX.And(boundY).ToBits() // panics: requires AVX512

		var results [16]int16
		boundX.And(boundY).ToInt16x16().Store(&results)
		for j, inside := range results {
			if inside != 0 {
				diffX, diffY := x-xPoints[i+j], y-yPoints[i+j]
				if diffX*diffX+diffY*diffY < distSq {
					log.Printf("Within distance: (%d, %d)\n", xPoints[i+j], yPoints[i+j])
				} else {
					log.Printf("In bounding box: (%d, %d)\n", xPoints[i+j], yPoints[i+j])
				}
			}
		}
	}
}
