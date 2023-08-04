package __

// Use the "Assem" build config (in GoLand) to generate assembly listing
// For a list of useful gcflags see ../../main.go:8 or select gcflags (filter by HELP category in the TODO window)
// ^\QNO_MATCH_just_for_assembly_listing\E$

import (
	"testing"
)

// TestGlobal is just here for assigning calculated values to.
// This is used in benchmarks to ensure that code is not optimised away, since if you assign a value to
// a variable visible outside the package the compiler doesn't know it's not used anywhere.
var TestGlobal any

func BenchmarkAssem1(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = i
	}
	TestGlobal = j
}

func BenchmarkAssem2(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = i * i
	}
	TestGlobal = j
}
