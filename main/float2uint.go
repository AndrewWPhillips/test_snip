package main

import (
	"fmt"
	"math"
)

// This demonstrates how float64 -> uint64 conversions are not monotonic as we approach 2^64
// I tracked this down to the "assembler" instruction CVTTSD2SQ

//func main() {
//	for f := float64(math.MaxUint64 - 5e4); f < float64(math.MaxUint64)+2e4; f += 1e4 {
//		fmt.Printf("%20.0f %v\n", f, uint64(f))
//		//_ = uint64(f)
//	}
//}

//func main() {
//	for p, n := 0.0, float64(math.MaxUint64-5e4); n < float64(math.MaxUint64)+2e4; p, n = n, n+1e4 {
//		println(uint64(p) <= uint64(n)) // should always print true (but doesn't)
//	}
//}

//func main() {
//	for ui64 := uint64(1<<64 - 2e4); ui64 != 0; ui64++ {
//		fmt.Println(float64(ui64))
//	}
//}

func main() {
	ui := uint64(math.MaxUint64)
	fmt.Println(math.MaxUint + 1)
	fmt.Println(ui)
	fmt.Println(ui + 1)
	math.Float64bits(0)
}
