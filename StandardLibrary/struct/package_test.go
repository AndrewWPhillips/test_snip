//go:build go1.23

package __

import (
	"log"
	"structs"
	"testing"
	"unsafe"
)

type (
	host struct {
		_ structs.HostLayout
		b byte
		f float64
	}
	noHost struct {
		b byte
		f float64
	}
	a struct {
		host
	}
	a1 struct {
		f noHost
	}
)

func TestHostLayout(t *testing.T) {
	var a a
	var a1 a1
	log.Print(unsafe.Sizeof(a))
	log.Print(unsafe.Sizeof(a1))
	//log.Print(unsafe.Offsetof(a.i))
	//log.Print(unsafe.Offsetof(a1.i))
}
