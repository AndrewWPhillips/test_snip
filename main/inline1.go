package main

// Just a test of compilation eg:
// $ go build -gcflags -m inline1.go

import (
	"log"
)

func main() {
	b()
	b1()
	log.Println(a(1, 2), a1(2, 3))
}

func a(i, j int) int {
	return i + j
}

func b() {
}

//go:noinline
func a1(i, j int) int {
	return i + j
}

//go:noinline
func b1() {
}
