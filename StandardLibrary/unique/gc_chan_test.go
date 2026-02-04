package __

import (
	"runtime"
	"testing"
	"time"
	"unique"
)

func TestUniqueHoldChanField(t *testing.T) {
	a := [8192]chan int{make(chan int)} //...
	type number struct {
		name string
		p    chan int
	}
	h := unique.Make(number{name: "four", p: a[0]})
	runtime.AddCleanup(&a, func(i int) { println("GC of a") }, 0)
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(h) // ensure handle is not GC'd
}
