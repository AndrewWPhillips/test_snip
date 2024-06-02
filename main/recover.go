package main

import (
	"fmt"
	"runtime/debug"
)

// Demonstration of recovering from diff types of panics for Alistair
func main() {
	testDiv0()
	testCustom()

	fmt.Println("Normal termination")
}

func testDiv0() int {
	defer func() {
		fmt.Println(recover())
	}()

	z := 0
	return 1 / z
}

func testCustom() int {
	defer func() {
		fmt.Println(recover())
		debug.PrintStack()
	}()

	panic("My own panic")
}
