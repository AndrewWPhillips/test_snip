package main

import (
	"fmt"
)

func main() {
	fmt.Println(f(), g())
}

func f() int {
	var r int
	defer func() {
		r = 42
	}()
	return r
}

func g() (r int) {
	defer func() {
		r = 42
	}()
	return r
}
