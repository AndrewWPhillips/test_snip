//go:build go1.22

// range func in for loops requires go 1.22 and GOEXPERIMENT=rangefunc

package __test

import (
	"fmt"
	"testing"
)

// BenchmarkOldRange is just for comparison - uses simple loop (no range func)
func BenchmarkOldRange(b *testing.B) { // 430 ns/op
	saved := 0
	for i := 0; i < b.N; i++ {
		for j := 0; j <= 999; j++ {
			saved = j
		}
	}
	fmt.Println(saved)
}

// BenchmarkBareRangeFunc uses the simplest range func
func BenchmarkBareRangeFunc(b *testing.B) { // 2300 ns/op
	saved := 0
	for i := 0; i < b.N; i++ {
		j := 0
		for range BareRange() {
			if j++; j > 999 {
				break
			}
			saved = j
		}
	}
	fmt.Println(saved)
}

func BareRange() func(func() bool) {
	return func(yield func() bool) {
		for {
			if !yield() {
				return
			}
		}
	}
}

// BenchmarkSingleRangeFunc is the simplest useful range func - uses a single parameter
func BenchmarkSingleRangeFunc(b *testing.B) { // 480 ns/op
	saved := 0
	for i := 0; i < b.N; i++ {
		for j := range SingleRange(999) {
			saved = j
		}
	}
	fmt.Println(saved)
}

func SingleRange(m int) func(func(int) bool) {
	return func(yield func(m int) bool) {
		for idx := 0; idx < m; idx++ {
			if !yield(idx) {
				return
			}
		}
	}
}

// BenchmarkSingleNoRangeFunc is equivalent to BenchmarkSingleRangeFunc (above)
func BenchmarkSingleNoRangeFunc(b *testing.B) { // 460 ns/op
	saved := 0
	for i := 0; i < b.N; i++ {
		for idx := 0; idx < 999; idx++ {
			saved = idx
		}
	}
	fmt.Println(saved)
}

var testSlice = []string{
	"first", "second", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "last",
}

func BenchmarkBackward1(b *testing.B) {
	jj, ss := 0, ""
	for i := 0; i < b.N; i++ {
		for j, s := range Backward2(testSlice) {
			jj, ss = j, s
		}
	}
	fmt.Println(jj, ss)
}

func Backward1(s []string) func(func(int, string) bool) { // ~20ns
	return func(yield func(int, string) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func Backward2(s []string) func(func(int, string) bool) { // ~16ns
	return func(yield func(int, string) bool) {
		for i := len(s); i > 0; i-- {
			if !yield(i-1, s[i-1]) {
				return
			}
		}
	}
}
