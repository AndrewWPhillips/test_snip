package main

import (
	"fmt"
	"math"
)

func main() {
	for f := float64(math.MaxUint64 - 5e4); f < float64(math.MaxUint64)+2e4; f += 1e4 {
		fmt.Printf("%20.0f %v %v\n", f, int64(f), uint64(f))
		//_ = uint64(f)
	}
}
