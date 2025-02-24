//go:build go1.22

package __

import (
	"fmt"
	"runtime"
	"testing"
)

func Test1_22Finalizer(t *testing.T) {
	v := new(int)
	*v = 42

	runtime.SetFinalizer(v, func(i *int) {
		fmt.Println("GC:", *i)
	})
	v = nil

	runtime.GC()
	runtime.GC()

	//var stats debug.GCStats
	//debug.ReadGCStats(&stats)
	//
	//var ms runtime.MemStats
	//runtime.ReadMemStats(&ms)
	//fmt.Printf("  GCStats:\n%v\n\n  MemStats:\n%v\n", stats, ms)
}
