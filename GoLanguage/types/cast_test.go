package __

import (
	"log"
	"testing"
)

func TestIntToBool(t *testing.T) {
	var b bool
	c := 1
	//b = bool(c)
	b = c != 0
	log.Println(b)
}

type myString string
type myByte byte
type mySlice []byte

func TestStringToByteSlice(t *testing.T) {
	var a myString = "a"
	var b string = "b"
	log.Println([]byte(a))
	log.Println([]byte(b))
	log.Println([]myByte(a))
	log.Println([]myByte(b))
	log.Println(mySlice(a))
	log.Println(mySlice(b))
}

func TestByteSliceToString(t *testing.T) {
	a := []myByte{65}
	b := mySlice{66}
	log.Println(string(a)) // OK in 1.18 but not in 1.15
	log.Println(string(b))
	log.Println(myString(a))
	log.Println(myString(b))
}
