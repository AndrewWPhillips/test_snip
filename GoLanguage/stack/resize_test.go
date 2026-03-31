package __test

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func BenchmarkStackAddress(b *testing.B) {
	b.ReportAllocs()
	a := 42
	p := &a
	for i := 0; i < b.N; i++ {
		a = r(&a, 100)
	}
	println(a, &a, p)
}

//go:noinline
func r(p *int, level int) int {
	if level == 0 {
		return *p
	}

	a := [ASize]int{}
	for i := 0; i < ASize; i++ {
		a[i] = rand.IntN(10) + 1
	}

	return r(p, level-1) + a[0]
}

func TestStackAddress(t *testing.T) {
	a := rand.Int()
	b := new(int)

	fmt.Println(a, *b)
}

func TestStackIncrease(t *testing.T) {
}

/*
func r(p *int, level int) *float64 {
	f := float64(3.14) * float64(level)
	if level%10000 == 0 {
		fmt.Printf("%5d: %p\n", level, &f)
		fmt.Printf("p: %p\n", p)
	}
	if level == 0 {
		return &f
	}

	return r(p, level-1)
}

func d(p *int) {
	var big struct {
		a [64]int
		b int
	}
	fmt.Printf("%p\n", &big.b)
	fmt.Printf("p: %p\n", p)
}

func b(p *int) {
	var big [2]int
	var big2 [2]int
	var big3 [2]int
	var big4 [2]int
	var big5 [2]int
	var big6 [2]int
	var big7 [2]int
	var big8 [2]int
	var big9 [2]int
	var biga [2]int
	var bigb [2]int
	var bigc [2]int
	var bigd [2]int

	fmt.Printf("%p\n", &big[0])
	fmt.Printf("%p\n", &big2[0])
	fmt.Printf("%p\n", &big3[0])
	fmt.Printf("%p\n", &big4[0])
	fmt.Printf("%p\n", &big5[0])
	fmt.Printf("%p\n", &big6[0])
	fmt.Printf("%p\n", &big7[0])
	fmt.Printf("%p\n", &big8[0])
	fmt.Printf("%p\n", &big9[0])
	fmt.Printf("%p\n", &biga[0])
	fmt.Printf("%p\n", &bigb[0])
	fmt.Printf("%p\n", &bigc[0])
	fmt.Printf("%p\n", &bigd[0])
	fmt.Printf("p: %p\n", p)
}
*/
