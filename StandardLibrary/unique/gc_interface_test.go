package __

import (
	"os"
	"runtime"
	"testing"
	"time"
	"unique"
)

// TestUniqueHoldInterfaceField is like TestUniqueHoldPointerField but with an interface
func TestUniqueHoldInterfaceField(t *testing.T) {
	a := new([8192]int)
	//h := unique.Make(struct{ any }{&a[42]}) // has problem but it's a ptr
	h := unique.Make(struct{ any }{a[42]}) // does not show the problem

	runtime.AddCleanup(a, func(i int) { println("GC of a") }, 0)
	a = nil // ensure a is not referenced - should be GC'd
	println("before GC")
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(h) // ensure handle is not GC'd
}

func TestUniqueHoldInterfaceField2(t *testing.T) {
	a := new([8192]struct {
		any
	})
	a[42].any = "abc"
	h := unique.Make(struct{ any }{a[42]})

	runtime.AddCleanup(a, func(i int) { println("GC of a") }, 0)
	a = nil // ensure a is not referenced - should be GC'd
	println("before GC")
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(h) // ensure handle is not GC'd

	os.Exit(0)
}

// The following are about testing when the values in interfaces get put on the heap

type IUneek interface{ F() }

type Uneek struct{ s string }

//go:noinline
func (u Uneek) F() {}

// BenchmarkUneek is for checking allocs/op (not nsecs/op)
func BenchmarkUneek(b *testing.B) { // 0 allocs
	a := Uneek{"uneek"}
	for b.Loop() {
		a.F()
	}
}

// BenchmarkIUneek is used to measure the number of allocs - shows that IUneek(a) escapes
func BenchmarkIUneek(b *testing.B) { // 1 alloc (when Uneek has a non-empty string field)
	//var a Uneek // 0 allocs
	a := Uneek{"uneek"} // 1 allocs
	for b.Loop() {
		IUneek(a).F()
	}
}

func TestUniqueHoldInterface(t *testing.T) {
	a := new([8192]Uneek)
	h := unique.Make(struct{ IUneek }{a[42]})
	runtime.AddCleanup(a, func(i int) { println("GC of a") }, 0)
	a = nil // ensure a is not referenced - should be GC'd
	println("before GC")
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(h) // ensure handle is not GC'd
}
