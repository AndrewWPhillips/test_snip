package __

import (
	"fmt"
	"runtime"
	"testing"
	"time"
	"unique"
)

// TestUniqueHoldPointer shows that passing a pointer into an array to unique prevents the array being GC'd
func TestUniqueHoldPointer(t *testing.T) {
	a := new([8192]int)
	h := unique.Make(&a[42])

	runtime.AddCleanup(a, func(i int) { println("GC of a") }, 0)
	// at this point the array 'a' is no longer itself referenced - should be available for GC
	println("before GC")
	runtime.GC()
	time.Sleep(time.Second) // wait for GC to finish
	println("after GC")
	runtime.KeepAlive(h) // ensure handle is not GC'd
}

// TestUniqueHoldPointerField is like TestUniqueHoldPointer but the pointer is in a struct field
func TestUniqueHoldPointerField(t *testing.T) {
	a := new([8192]int)
	h := unique.Make(struct{ *int }{&a[42]})

	runtime.AddCleanup(a, func(i int) { println("GC of a") }, 0)
	a = nil // ensure a is not referenced - should be GC'd
	println("before GC")
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(h) // ensure handle is not GC'd
}

func TestUniqueHoldPointer3(t *testing.T) {
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

func TestUniqueHoldPointerField2(t *testing.T) {
	a := [8192]int{0, 2, 3, 4, 5, 6, 7, 8, 9} //...
	type number struct {
		name string
		p    *int
	}
	h := unique.Make(number{name: "four", p: &a[4]})
	runtime.AddCleanup(&a, func(i int) { println("GC of a") }, 0)
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(h) // ensure handle is not GC'd
}
