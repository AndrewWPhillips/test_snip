//go:build go1.22

// range func in for loops requires go 1.22 and GOEXPERIMENT=rangefunc

package __

import (
	"fmt"
	"strconv"
	"testing"
)

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

// OneRange is a  range func yielding one value (int) ------------
func OneRange() func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func TestOneRangeFunc(t *testing.T) {
	for i := range OneRange() {
		println(i)
	}
}

func TwoRange() func(func(int, int) bool) {
	return func(yield func(i, j int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, i*i) {
				return
			}
		}
	}
}

func TestTwoRangeFunc(t *testing.T) {
	for i, j := range TwoRange() {
		println(i, j)
	}
}

/*
func TestTwoRangeOneFunc(t *testing.T) {
	for i := range TwoRange() { // range over TwoRange() (value of type func(func(int, int) bool)) must have two iteration variables
		println(i)
	}
}
*/

func StringRange() func(func(string) bool) {
	return func(yield func(s string) bool) {
		for i := 0; i < 10; i++ {
			if !yield(strconv.Itoa(i)) {
				return
			}
		}
	}
}

func TestStringRangeFunc(t *testing.T) {
	for s := range StringRange() {
		fmt.Printf("%q %T\n", s, s)
	}
}

/*
func ThreeRange() func(func(int, int, int) bool) {
	return func(yield func(i, j, k int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i, i*2, i*i) {
				return
			}
		}
	}
}

func TestThreeRangeFunc(t *testing.T) {
	for i, j, k := range ThreeRange() { // expected at most 2 expressions
		println(i, j, k)
	}
}
*/

// IntRange is a  range func yielding two values (int, int) ------------
func IntRange(m int) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		println("yield start")
		for i := 0; i < m; i++ {
			println("yield", i)
			if !yield(i, i) {
				println("yield break")
				return
			}
		}
		println("yield end")
	}
}

func TestRangeFunc(t *testing.T) {
	for i, x := range IntRange(15) {
		fmt.Printf("i: %d %T, x: %d, %T\n", i, i, x, x)
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

func TestNoBreakRangeFunc(t *testing.T) {
	for i, x := range NoBreakRange(15) {
		fmt.Printf("i: %d %T, x: %d, %T\n", i, i, x, x)
		if i > 5 {
			break
		}
	}
}

// Backward is a  range func yielding two values (int, string): the index and value from a slice in reverse order
func Backward(s []string) func(func(int, string) bool) {
	return func(yield func(int, string) bool) {
		for i := len(s); i > 0; i-- {
			if !yield(i-1, s[i-1]) {
				return
			}
		}
	}
}

func TestBackwardRangeFunc(t *testing.T) {
	s := []string{"hello", "world"}
	for i, x := range Backward(s) {
		fmt.Println(i, x)
	}
}
