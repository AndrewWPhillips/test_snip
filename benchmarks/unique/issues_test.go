package __

// test reported Go unique package issues

import (
	"fmt"
	"runtime"
	"testing"
	"unique"
)

// TestIssue71926 does not show the issue - we need to look at the asm
func TestIssue71926(t *testing.T) {
	h1 := unique.Make("abc")
	d := []byte("abc")
	h2 := unique.Make(string(d)) // look at code generated
	//h2 := unique.Make(unsafe.String(&d[0], len(d)))

	fmt.Println(h1, h2)
}

// TestUniqueEmpty is not really useful
func TestUniqueEmpty(t *testing.T) {
	var e1, e2 [0]int
	h1 := unique.Make(e1)
	h2 := unique.Make(e2)
	fmt.Println(h1, h2)
}

// TestUniqueSubString - check if interning a substring prevents the string from being garbage collected
func TestUniqueSubString(t *testing.T) {
	// Create a long string on the heap
	s := "gfdpjgfdgfdogfodijgfdogfdoijgfdojgfdogfdjgfodijgfodjgfdoojgfdoijgfdoi"
	//runtime.AddCleanup(&s, func(s string) {
	//	fmt.Println("cleanup string")
	//}, s)
	//
	runtime.SetFinalizer(&s, func(s *string) {
		fmt.Println("cleanup string")
	})

	h1 := unique.Make(s[1:2])
	//runtime.AddCleanup(&h1, func(s unique.Handle[string]) {
	//	fmt.Println("cleanup handle")
	//}, h1)
	runtime.SetFinalizer(&h1, func(h *unique.Handle[string]) {
		fmt.Println("cleanup handle")
	})

	s = ""                       // remove reference to the original string (previous vales of s)
	h1 = unique.Handle[string]{} // remove ref to handle (except KeepAlive below)

	// This should cleanup the original string unless unique.Make kept a pointer into it
	fmt.Println("before GC 1")
	runtime.GC()
	fmt.Println("after GC 1")

	runtime.KeepAlive(h1) // remove ref to handle

	// This should cleanup the handle
	fmt.Println("before GC 2")
	runtime.GC()
	fmt.Println("after GC 2")
}
