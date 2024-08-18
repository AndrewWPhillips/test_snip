package __

import (
	"fmt"
	"testing"
)

// Tests (benchmarks) for possible meetup quiz on performance

// Q1 about map key of a concatenated string vs struct of 2 strings

// storeStr stores int values in a map with a string key
var storeStr = map[string]int{}

func put(k string, v int) {
	storeStr[k] = v
}

func get(k string) int {
	return storeStr[k]
}

func put2(k1, k2 string, v int) {
	storeStr[k1+":"+k2] = v
}

func get2(k1, k2 string) int {
	return storeStr[k1+":"+k2]
}

// storeStruct stores int values in a map a "two string" key
var storeStruct = map[struct {
	k1 string
	k2 string
}]int{}

func put3(k1, k2 string, v int) {
	storeStruct[struct {
		k1 string
		k2 string
	}{k1, k2}] = v
}

func get3(k1, k2 string) int {
	return storeStruct[struct {
		k1 string
		k2 string
	}{k1, k2}]
}

func TestStore(t *testing.T) {
	put("four", 2)
	fmt.Println(storeStr)
	storeStr = make(map[string]int)
	put2("four", "two", 42)
	fmt.Println(storeStr)
	put3("four", "two", 42)
	fmt.Println(storeStruct)
}

func BenchmarkCombinedString(b *testing.B) { // 1 alloc, 70ns
	lhs, rhs := "left", "right"
	const compare = "left:right"
	combined := ""
	result := false
	for i := 0; i < b.N; i++ {
		combined = lhs + ":" + rhs
		result = compare == combined
	}
	fmt.Println(result)
}

func BenchmarkCombinedString2(b *testing.B) { // 0 allocs!, 27ns
	lhs, rhs := "left", "right"
	const compare = "left:right"
	result := false
	for i := 0; i < b.N; i++ {
		result = compare == lhs+":"+rhs
	}
	fmt.Println(result)
}

func BenchmarkStringPair(b *testing.B) { // 0 allocs, <1ns
	compare := struct{ a, b string }{"left", "right"}
	pair := compare
	lhs, rhs := "left", "right"
	result := false
	for i := 0; i < b.N; i++ {
		pair.a, pair.b = lhs, rhs
		result = pair == compare
	}
	fmt.Println(result)
}

// Q2 avoid copying during loop iteration

func TestLoop(t *testing.T) {
	a := [1024]int{0, 1, 2}
	sa := make([][1024]int, 0, 256)

	for range 256 {
		sa = append(sa, a)
	}

	// sa is 2Mb
	var ii, vv int
	for i, v := range sa {
		ii = i
		vv = len(v)
	}
	println(ii, vv)
}

const (
	SliceLen    = 1024
	ElementSize = 4096
)

type T [ElementSize]byte

var a = T{0, 1, 2} // 4K byte array

func BenchmarkLoopCopy(b *testing.B) { // 430 ns/op
	// Make a 4MB slice (1024 x 4KB)
	sa := make([]T, 0, SliceLen)
	for i := 0; i < SliceLen; i++ {
		sa = append(sa, a)
	}

	var ii, vv int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for idx, value := range sa {
			ii = idx
			vv = len(value)
		}
	}
	fmt.Println(ii, vv)
}

func BenchmarkLoopNoCopy(b *testing.B) { // 413
	// Make a 4MB slice (1024 x 4KB)
	sa := make([]T, 0, SliceLen)
	for i := 0; i < SliceLen; i++ {
		sa = append(sa, a)
	}

	var ii, vv int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for idx := range sa {
			ii = idx
			vv = len(sa[idx])
		}
	}
	fmt.Println(ii, vv)
}

// Q3 - removing duplicates - see benchmarks/slice/uniq_test.go:43
