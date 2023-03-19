package __

import (
	"fmt"
	"log"
	"testing"
	"unsafe"
)

func TestUintptr(t *testing.T) {
	i := 3
	p := &i
	log.Println(p, *p) // 0xc000120438 3
	uip := unsafe.Pointer(p)
	p = nil
	p = (*int)(unsafe.Pointer(uip))
	log.Println(p, *p) // 0xc000120438 3

	c := make(chan int)
	log.Printf("%v %v %v %x\n",
		c, &c, unsafe.Pointer(&c), uintptr(unsafe.Pointer(&c))) // 0xc0001081e0 0xc000146028 0xc000146028 c000146028

	s := fmt.Sprint(c)
	log.Printf("%v %T", s, s) // 0xc0001081e0 string
}
