package __

import (
	"log"
	"testing"
	"unsafe"
)

func TestSizeof(t *testing.T) {
	var a interface{}
	log.Println(unsafe.Sizeof(struct{}{}))
	log.Println(unsafe.Sizeof(a))
}
