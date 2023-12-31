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
	f := float64(math.MaxUint64 - 2000)
	log.Println(uint64(f) >= uint64(f-1e4)) // false, expected true
	log.Println(uint64(f), uint64(f-1e4), uint64(f-10000), uint64(math.MaxUint64))
}

// floating point values larger than max uint (and slightly less) convert to 9,223,372,036,854,775,808 (2^63)
// but should become 18,446,744,073,709,551,615 (2^64-1)
func TestFToU64A(t *testing.T) {
	f := float64(math.MaxUint64)
	log.Println(f, uint64(f)) // compiles OK but prints 9223372036854775808

	f = float64(math.MaxUint64 - 2e4)
	log.Println(f, uint64(f)) // 18446744073709531136

	f = float64(math.MaxUint64)
	f *= 2
	log.Println(f, uint64(f)) // 9223372036854775808

	//log.Println(uint64(float64(math.MaxUint64))) // cannot convert float64(math.MaxUint64) (constant 18446744073709551616 of type float64) to type uint64
	log.Println(uint64(float64(math.MaxUint64 - 2e4)))
}

func TestFToU64B(t *testing.T) {
	//start := uint64(9_223_372_036_854_775_800)
	var start uint64 = math.MaxUint64
	for i := start; i > (math.MaxUint64 - 2000); i-- {
		log.Printf("%d %20f\n", i, float64(i))
	}
}
