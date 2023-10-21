package main

import (
	"log"
	"math/rand"
)

func main() {
	log.Println(c(), d(1))
}

func c() int {
	return d(0)
}

//go:noinline  // prevent inlining
func d(i int) int {
	var x [1000]int
	x[i] = rand.Int()
	return x[i]
}
