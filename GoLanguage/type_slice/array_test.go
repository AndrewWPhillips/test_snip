package __

import (
	"log"
	"testing"
	"unsafe"
)

func TestArrayToSlice(t *testing.T) {
	var a [20]int
	log.Println("Size of a:", unsafe.Sizeof(a)) // 160

	b := a
	log.Println("Size of b:", unsafe.Sizeof(b)) // 160

	c := a[:]
	log.Println("Size of c:", unsafe.Sizeof(c)) // 24

	func(d []int) {
		log.Println("Size of d:", unsafe.Sizeof(d)) // 24
	}(a[:]) // convert array to slice before passing to func expecting a slice
}
