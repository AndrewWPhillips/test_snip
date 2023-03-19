package __

import (
	"log"
	"math"
	"testing"
)

// TestFloat64ToUint64 demonstrates a "bug" where conversion is not monotonic close to max uint64 value
// Note that the conversion can't be exactly represented but as float64 value increases so should uint64
func TestFloat64ToUint64(t *testing.T) {
	//f := float64(18_446_744_073_709_551_615)  // 2^64-1 ~=> 1.844674407370955e19
	f := float64(math.MaxUint64)
	log.Println(uint64(f) >= uint64(f-1e4)) // false, expected true
	log.Println(uint64(f), uint64(f-1000), uint64(f-1001), uint64(f-1002), uint64(math.MaxUint64))
}
