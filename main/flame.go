package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	var i, s, f, sf int
	for i = 0; i < 1e8; i++ {
		s = sum(20)
		f = fact(20)
		sf = sumPlusFact(20)
	}
	fmt.Println(s, f, sf)
}

func sumPlusFact(n int) int {
	return sum(n) + fact(n)
}

func sum(n int) int {
	r := 0
	for i := 0; i <= n; i++ {
		r += i
	}
	return r
}

func fact(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}
