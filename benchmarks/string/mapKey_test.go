package __

import (
	"fmt"
	"testing"
)

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
