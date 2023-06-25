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
	fmt.Println([]byte(a))
	fmt.Println([]byte(b))
	fmt.Println([]myByte(a))
	fmt.Println([]myByte(b))
	fmt.Println(mySlice(a))
	fmt.Println(mySlice(b))
}

func TestByteSliceToString(t *testing.T) {
	a := []myByte{65}
	b := mySlice{66}
	fmt.Println(string(a)) // OK in 1.18 but not in 1.15
	fmt.Println(string(b))
	fmt.Println(myString(a))
	fmt.Println(myString(b))
}
