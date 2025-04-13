package __

import (
	"strings"
	"testing"
	"unsafe"
)

var sGlobal = "abc"

// TestStringData works out where in memory strings are stored by using
// unsafe.StringData to get a pointer to the underlying bytes of a string
func TestStringData(t *testing.T) {
	const sConst = "abc"
	sLoc1 := "abc"
	sLoc2 := "ab"
	sLoc3 := "ab" + "c"
	sSub1 := sLoc1[:]
	sSub2 := sLoc1[:2]   // ab
	sSub3 := sLoc1[1:2]  // b
	sSub4 := "abcde"[:3] // abc
	sCalc1 := sLoc1 + "x"
	sCalc2 := sSub2 + "c" // abc
	sBuilt1 := strings.Builder{}
	sBuilt1.WriteString("abc")

	println(unsafe.StringData("abc"), // Static A (0xff2f0b)
		unsafe.StringData(sGlobal),          // Static A
		unsafe.StringData(sConst),           // Static A
		unsafe.StringData(sLoc1),            // Static A
		unsafe.StringData(sLoc2),            // Static B (0xff2e48)
		unsafe.StringData(sLoc3),            // Static A
		unsafe.StringData(sSub1),            // Static A
		unsafe.StringData(sSub2),            // Static A
		unsafe.StringData(sSub3),            // Static A+1 (0xff2f0c)
		unsafe.StringData(sSub4),            // Static C (0xff3256)
		unsafe.StringData(sCalc1),           // Stack M (0xc000051ef8)
		unsafe.StringData(sCalc2),           // Stack N
		unsafe.StringData(sBuilt1.String()), // Stack P
	)
}
