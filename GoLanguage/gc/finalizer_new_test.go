//go:build go1.24

package __

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestCompilerBug(t *testing.T) {
	//var i8 int8
	//for i8 = range [1]bool{true} { // index var should allow any int type TODO: check spec
	//	println(i8)
	//}
}

func TestMultipleSmall(t *testing.T) {
	type smallInt int16
	type smallType struct {
		i     smallInt
		dummy [60]uint8 // used to make struct small or big
	}
	a := make([]*smallType, 20)

	// Add value and finalizers
	for i := range a {
		a[i] = new(smallType)
		a[i].i = 40 + smallInt(i)
		runtime.SetFinalizer(a[i], func(pi *smallType) {
			fmt.Println("Finalizer:", pi.i)
		})
	}

	runtime.GC()
	runtime.GC()
	time.Sleep(2 * time.Second)
	//runtime.KeepAlive(a[0])
}

func TestRuntimeAddCleanup(t *testing.T) {
	p := new(int)
	*p = 42
	//runtime.SetFinalizer(p, func(p2 *int) { println(*p2) })
	runtime.AddCleanup(p, func(v int) { println(v) }, *p)
	runtime.GC()
	runtime.GC()
}

func TestRuntimeAddCleanup2(t *testing.T) {
	p := new(struct {
		dummy [16]int8
		v     int
	})
	p.v = 42
	runtime.AddCleanup(&p.v, func(v int) { println(v) }, p.v)
	runtime.GC()
}

func TestAddCleanup(t *testing.T) {
	type T int64
	a := make([]*T, 20)

	// Add value and finalizers
	for i := range a {
		a[i] = new(T)
		*a[i] = 40 + T(i)
		//runtime.SetFinalizer(a[i], func(pi *T) {
		//	fmt.Println("Finalizer:", *pi)
		//})
		runtime.AddCleanup(a[i], func(v T) {
			fmt.Println("Cleanup:", v)
		}, *a[i])
	}

	runtime.GC()
	runtime.GC()
}
