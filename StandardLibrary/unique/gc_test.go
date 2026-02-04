package __

import (
	"fmt"
	"runtime"
	"testing"
	"time"
	"unique"
)

// TestUniqueHoldSubString checks that unique.Make() makes a copy of a string so that intern of
// a substring of a larger string does not prevent the larger string from being GCed
// See also gc_string_test.go
func TestUniqueHoldSubString(t *testing.T) {
	a := "abcdefhijk"
	h := unique.Make(a[1:2])

	runtime.AddCleanup(&a, func(i int) { println("GC of a") }, 0)
	a = "" // ensure the original string is no longer referenced
	println("before GC")
	runtime.GC()
	time.Sleep(time.Microsecond) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(h)
}

// TestUniqueHoldStringField is like TestUniqueHoldSubString but checks that a substring as
// a struct field does not prevent the larger string from being GCed
func TestUniqueHoldStringField(t *testing.T) {
	a := "abcdefhijk"
	unique.Make(struct{ string }{a[1:2]})

	runtime.AddCleanup(&a, func(i int) { println("GC of a") }, 0)
	a = "" // ensure the original string is no longer referenced
	println("before GC")
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
}

// TestUniqueHoldArrayElement shows that passing a pointer into an array to unique prevents the array being GC'd
func TestUniqueHoldArrayElement(t *testing.T) {
	a := new([8192]int)
	//unique.Make(&a[42])

	runtime.AddCleanup(a, func(i int) { println("GC of a") }, 0)

	a = nil // ensure a is no longer referenced
	runtime.GC()
	time.Sleep(time.Second) // wait for GC to finish
	println("after GC")
}

func TestUniqueHoldArrayElement2(t *testing.T) {
	a := new([8192]int)
	h := unique.Make(&a[42])

	// Stop using a and see if it is garbage collected
	runtime.AddCleanup(a, func(i int) { fmt.Println("GC of a", i) }, a[42])

	fmt.Println("stop using array")
	a = nil
	runtime.GC()
	runtime.KeepAlive(h)

	fmt.Println("stop using handle")
	h = unique.Handle[*int]{}
	runtime.GC()
	time.Sleep(time.Second)
	fmt.Println("END")
}

func TestUniqueHoldArrayElement3(t *testing.T) {
	a := new([8192]int)
	i := 42
	_ = i

	a[4242] = 42

	h := unique.Make(&a[4242])
	//h := unique.Make(&i)
	fmt.Println(h, *h.Value())

	// Stop using a and see if it is garbage collected
	runtime.AddCleanup(a, func(a [8192]int) { fmt.Println("GC of a") }, *a)

	fmt.Println("stop using array")
	a = nil
	runtime.GC()
	time.Sleep(time.Millisecond)
	fmt.Println("stop using handle")
	h = unique.Handle[*int]{}
	runtime.GC()
	time.Sleep(time.Millisecond)
	fmt.Println("END")
}
