package __

import (
	"fmt"
	"strings"
	"testing"
)

var TTT any // global used to make sure code is not optimized away

// BenchmarkConcat tests the performance of manual concatenation
// Go 1.15.5,  Windows 10, Xeon(R) CPU E5-1650 v4 @ 3.60GHz : 107ns/op
// Go 1.18.3,  Windows 10, Xeon(R) CPU E5-1650 v4 @ 3.60GHz : 105ns/op
// Go 1.17.10, Windows 10, AMD Ryzen 5 2600 6core @ 3.40GHz : 207ns/op
// Go 1.18.4,  Windows 10, AMD Ryzen 5 2600 6core @ 3.40GHz : 203ns/op
func BenchmarkConcat(b *testing.B) { // ~100 ns, 2 allocs/op
	var result string
	ss := []string{"sadsadsa", "dsadsakdas;k", "8930984"}
	for i := 0; i < b.N; i++ {
		var s, sep string
		for j := 0; j < len(ss); j++ {
			s += sep + ss[j]
			sep = " "
		}
		result = s
	}
	TTT = result
}

// BenchmarkJoin tests the performance of strings.Join
// Go 1.15.5,  Windows 10, Xeon(R) CPU E5-1650 v4 @ 3.60GHz : 57ns/op
// Go 1.18.3,  Windows 10, Xeon(R) CPU E5-1650 v4 @ 3.60GHz : 53ns/op
// Go 1.17.10, Windows 10, AMD Ryzen 5 2600 6core @ 3.40GHz : 1030ns/op ********
// Go 1.18.4,  Windows 10, AMD Ryzen 5 2600 6core @ 3.40GHz : 1031ns/op
func BenchmarkJoin(b *testing.B) {
	var result string
	ss := []string{"sadsadsa", "dsadsakdas;k", "8930984"}
	for i := 0; i < b.N; i++ {
		result = strings.Join(ss, " ")
	}
	TTT = result
}

// BenchmarkBuilder tests the performance of strings.Builder
// Go 1.15.5,  Windows 10, Xeon(R) CPU E5-1650 v4 @ 3.60GHz : 54ns/op
// Go 1.18.3,  Windows 10, Xeon(R) CPU E5-1650 v4 @ 3.60GHz : 50ns/op
// Go 1.17.10, Windows 10, AMD Ryzen 5 2600 6core @ 3.40GHz : 947ns/op ********
// Go 1.18.4,  Windows 10, AMD Ryzen 5 2600 6core @ 3.40GHz : 942 ns/op
func BenchmarkBuilder(b *testing.B) { // 1 alloc/op
	var result string
	ss := []string{"sadsadsa", "dsadsakdas;k", "8930984"}
	for i := 0; i < b.N; i++ {
		var s strings.Builder
		// Grow builder to expected max length (maybe this
		// needs to be calculated dep. on your requirements)
		s.Grow(32)
		var sep string
		for j := 0; j < len(ss); j++ {
			s.WriteString(ss[j])
			s.WriteString(sep)
			sep = " "
		}
		result = s.String()
	}
	TTT = result
}

func BenchmarkSprintf(b *testing.B) { // ~200 ns, 4 allocs/op
	var result string
	ss := []string{"sadsadsa", "dsadsakdas;k", "8930984"}
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s %s %s", ss[0], ss[1], ss[2])
		result = s
	}
	TTT = result
}
