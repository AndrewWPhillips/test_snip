package __test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

/*
//This is from "Anatomy of Go" book, but does not work (at least with Go 1.25)

func TestGrowStack(t *testing.T) {
	x := 1
	xUnsafePtr := unsafe.Pointer(&x)
	xUintptr := uintptr(xUnsafePtr)
	println("Before: xPtr", &x, "xUnsafePtr", xUnsafePtr, "xUintptr", xUintptr)
	growStack()
	println("After:  xPtr", &x, "xUnsafePtr", xUnsafePtr, "xUintptr", xUintptr)
}

func growStack() {
	x := [1024]byte{}
	_ = fmt.Sprint(x)
}
*/

func TestGrowStack(t *testing.T) {
	x := 1
	addressBefore := uintptr(unsafe.Pointer(&x))
	growRecurse(1024)
	addressAfter := uintptr(unsafe.Pointer(&x))
	if addressBefore != addressAfter {
		fmt.Printf("Before: %[1]x %[1]d\nAfter:  %[2]x %[2]d\n", addressBefore, addressAfter)
	}
}

func growRecurse(level int) {
	if level == 0 {
		return
	}
	//if level%100 == 0 {
	//	fmt.Printf("address: %x\n", uintptr(unsafe.Pointer(&level)))
	//}
	growRecurse(level - 1)
}

func TestPin(t *testing.T) {
	pinner := runtime.Pinner{}
	_ = pinner
}

func TestFunctionPointer(t *testing.T) {
	println(reflect.ValueOf(TestPin).Pointer())
	println(reflect.ValueOf(TestFunctionPointer).Pointer())
}
