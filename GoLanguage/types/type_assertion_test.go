package __

import (
	"bytes"
	"fmt"
	"log"
	"testing"
	"unsafe"
)

func TestTypeAssertion(t *testing.T) {
	var x interface{}
	var s fmt.Stringer
	a := 1
	b := bytes.NewBufferString("abc")

	x = a
	{
		v, ok := x.(int)
		log.Printf("%T %v %v\n", v, v, ok) // int 1 true
	}

	x = b
	{
		v, ok := x.(*bytes.Buffer)
		log.Printf("%T %v %v\n", v, v, ok) // *bytes.Buffer abc true
	}

	x = s
	{
		v, ok := x.(fmt.Stringer)
		log.Printf("%T %v %v\n", v, v, ok) // nil nil false
	}

	s = b
	x = s
	{
		v, ok := x.(fmt.Stringer)
		log.Printf("%T %d %v %v\n", v, unsafe.Sizeof(v), v, ok) // *bytes.Buffer 16 abc true (Printf %T and %v looks inside the interface)
	}
	{
		v, ok := x.(interface{})
		log.Printf("%T %d %v %v\n", v, unsafe.Sizeof(v), v, ok) // *bytes.Buffer 16 abc true
	}
	{
		v, ok := x.(*bytes.Buffer)
		log.Printf("%T %d %v %v\n", v, unsafe.Sizeof(v), v, ok) // *bytes.Buffer 8 abc true
	}
	{
		v, ok := x.(int)
		log.Printf("%T %d %v %v\n", v, unsafe.Sizeof(v), v, ok) // int 8 0 false
	}
}
