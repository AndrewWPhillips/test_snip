package main

import (
	"fmt"
	"github.com/pkg/profile"
)

// This is the code used with my blog to generate flame graphs
// https://andrewwphillips.github.io/blog/flame-graphs.html

// $ go run flame_graph.go
// 2024/01/15 23:03:59 profile: cpu profiling enabled, cpu.pprof
// 210 2432902008176640000 2432902008176640210
// 2024/01/15 23:04:02 profile: cpu profiling disabled, cpu.pprof
//
// $ go tool pprof -http :6060 cpu.pprof
// Serving web UI on http://localhost:6060

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	var i, s, f, sf int
	for i = 0; i < 2e8; i++ {
		s = sum(20)
		f = fact(20)
		sf = sumPlusFact(20)
	}
	fmt.Println(s, f, sf)
}

func sumPlusFact(n int) int {
	return sum(n) + fact(n)
}

//func sum(n int) int {
//	r := 0
//	for i := 0; i <= n; i++ {
//		r += i
//	}
//	return r
//}

func sum(n int) int {
	return (n * (n + 1)) / 2
}

func fact(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}
