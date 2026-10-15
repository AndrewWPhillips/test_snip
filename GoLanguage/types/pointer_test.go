package __

import (
	"testing"
	"unsafe"
)

func TestUnsafePointer(t *testing.T) {
	var p *int
	var up unsafe.Pointer
	var uip uintptr

	up = unsafe.Pointer(p)
	up = unsafe.Pointer(&p)
	up = unsafe.Pointer(&up)

	// uip = uintptr(p) // cannot convert an expression of *int to uintptr
	uip = uintptr(up) // OK
	println(p, up, uip)
}

func TestPointerOrder(t *testing.T) {
	i := 1
	p1 := &i
	var p2 *int

	if p1 != p2 {
		println("Can compare pointers")
	}

	//if p1 < p2 {                                 // invalid operation: < is not allowed on pointers
	//if unsafe.Pointer(p1) < unsafe.Pointer(p2) { // invalid operation: unsafe.Pointer(p1) < unsafe.Pointer(p2) (operator < not defined on unsafe.Pointer)
	if uintptr(unsafe.Pointer(p1)) < uintptr(unsafe.Pointer(p2)) { // OK
		println("uintptr is orderable")
	}
}
