//go:build go1.23

// Tests of range funcs used in for-range statments
// Originally these tests were for go 1.22 with GOEXPERIMENT=rangefunc
// Go 1.23 added these calling them iterators with support in the iter package
// See also
//	benchmarks/iter/iter_test.go:1
//	StandardLibrary/iter/seq0_test.go:1

package __

import (
	"fmt"
	"iter"
	"math"
	"testing"
	"time"
)

////////////// RANGE FUNCS ///////////////

// NoRange is a range func yielding no values ---------------
func NoRange() func(func() bool) {
	return func(yield func() bool) {
		for {
			if !yield() {
				return
			}
		}
	}
}

// OneRange is a range func yielding one value (int) ------------
func OneRange() iter.Seq[int] {
	//func OneRange() func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// TwoRange yields a sequence of two ints
// Parameters
//
//	m = length of the range
//	show = print info about control flow in the range func
//
// func TwoRange(m int, show bool) func(func(int, float64) bool) {
func TwoRange(m int, show bool) iter.Seq2[int, float64] {
	return func(yield func(int, float64) bool) {
		if show {
			println("yield start")
		}
		for i := 0; i < m; i++ {
			if show {
				println("yield", i)
			}
			if !yield(i, math.Sqrt(float64(i))) {
				if show {
					println("yield break")
				}
				return
			}
		}
		if show {
			println("yield end")
		}
	}
}

// ThreeRange yields a sequence of three ints
// but if you try to use it with a for-range statement you get ERROR: yield func has too many parameters
func ThreeRange() func(func(int, int, int) bool) {
	return func(yield func(i, j, k int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, i*2, i*i) {
				return
			}
		}
	}
}

// NoBreakRange is a range func that breaks the rules - should return when yield() returns false
func NoBreakRange(m int) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		for i := 0; i < m; i++ {
			yield(i, i) // should return on false
			// The break in the loop gives: panic: runtime error: range function continued iteration after exit [recovered]
		}
	}
}

// BackwardStr is a range func yielding two values (int, string): the index and value from a slice in reverse order
func BackwardStr(s []string) func(func(int, string) bool) {
	return func(yield func(int, string) bool) {
		for i := len(s); i > 0; i-- {
			if !yield(i-1, s[i-1]) {
				return
			}
		}
	}
}

// Backward is a generic range func yielding two values (int, T): the index and value from a slice in reverse order
// func Backward[T any](s []T) func(func(int, T) bool) {
func Backward[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := len(s); i > 0; i-- {
			if !yield(i-1, s[i-1]) {
				return
			}
		}
	}
}

func blogIterator(yieldFunc func(int, float64) bool) {
	for i := 0; i < 7; i++ {
		if !yieldFunc(i, math.Sqrt(float64(i))) {
			return
		}
	}
	return
}

/////////////// TESTS ///////////////

func TestSimpleNoRangeFunc(t *testing.T) {
	for range NoRange() {
	}
}

func TestNoRangeFuncBreak(t *testing.T) {
	i := 0
	for range NoRange() {
		if i++; i > 10 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
	fmt.Println("end")
}

func TestNoRangeFuncReturn(t *testing.T) {
	defer fmt.Println("returned")
	i := 0
	for range NoRange() {
		if i++; i > 10 {
			fmt.Println("return")
			return
		}
		fmt.Println(i)
	}
	fmt.Println("end")
}

// TestOneRangeFunc0 ranges over OneRange but without any iteration variable
func TestOneRangeFunc0(t *testing.T) {
	for range OneRange() { // OK Go 1.23 but gopls (1.22?) gives ERROR: requires exactly one iteration variable
		println(0)
	}
}

func TestOneA(t *testing.T) {
	for i := range OneRange() {
		println(i)
	}
}

func TestOneANoForRange(t *testing.T) {
	it := OneRange()
	it(func(i int) bool {
		println(i)
		return true
	})
}

func TestOneB(t *testing.T) {
	for i := range OneRange() {
		if i == 3 {
			continue
		}
		println(i)
		if i > 6 {
			goto endLoop // or break or return
		}
	}
endLoop:
}

func TestOneBNoForRange(t *testing.T) {
	it := OneRange()
	it(func(i int) bool {
		if i == 3 {
			return true // do nothing further but continue looping
		}
		println(i)
		if i > 6 {
			return false // stop looping
		}
		return true
	})
}

func TestTwoRangeFunc(t *testing.T) {
	// for i := range TwoRange(5, false) { // OK but gopls gives ERROR: must have two iteration variables
	for i, j := range TwoRange(5, false) {
		fmt.Printf("i: %d %T, j: %v, %T\n", i, i, j, j)
	}
}

func TestTwoRangeShow(t *testing.T) {
	println("start")
	for i, j := range TwoRange(3, true) {
		fmt.Println(i, j)
		if i > 1 {
			break
		}
	}
	println("end")
}

func TestNoBreakRangeFunc(t *testing.T) {
	for i, x := range NoBreakRange(15) {
		fmt.Printf("i: %d %T, x: %d, %T\n", i, i, x, x)
		if i > 5 {
			break
		}
	}
}

func TestBackwardRangeFunc(t *testing.T) {
	s := []string{"hello", "there", "world"}
	for i, x := range BackwardStr(s) {
		fmt.Println(i, x)
	}
	for i, x := range Backward(s) {
		fmt.Println(i, x)
	}
}

func TestBlog(t *testing.T) {
	for v, root := range blogIterator {
		if root == 0.0 {
			continue
		}
		if root >= 2.0 {
			break
		}
		println(v, root)
	}
	time.NewTicker()
}
