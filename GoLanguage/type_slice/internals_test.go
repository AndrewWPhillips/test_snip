package __

import (
	"testing"
	"unsafe"
)

func TestSliceInternals(t *testing.T) {
	mySlice := make([]byte, 3, 5)

	//sliceArray := uintptr(unsafe.Pointer(&mySlice))
	sliceLength := unsafe.Pointer(uintptr(unsafe.Pointer(&mySlice)) + uintptr(8))
	sliceCapacity := unsafe.Pointer(uintptr(unsafe.Pointer(&mySlice)) + uintptr(16))
	println("array:", *(*unsafe.Pointer)(unsafe.Pointer(&mySlice)))
	println("length:", *(*int)(sliceLength))
	println("capacity:", *(*int)(sliceCapacity))
	println("mySlice:", mySlice)
}
