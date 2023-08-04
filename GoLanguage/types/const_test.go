package __

import (
	"fmt"
	"testing"
)

// TestBuiltinFunctionType tests the type of expression only involving builtin function and constants
func TestBuiltinFunctionType(t *testing.T) {
	a := [...]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var x, y byte
	x = 1 << len(a) / 128    // len(a) is a const => untyped number
	y = 1 << len(a[:]) / 128 // 1 << len(a[:]) appears to return byte not int
	fmt.Println(x, y)        // 4 0
}

func TestBuiltinFunctionType1(t *testing.T) {
	a := [9]byte{}
	var x, y byte
	x = 1 << len(a) / 128
	y = 1 << len(a[:]) / 128
	fmt.Println(x, y) // 4 0
}

func TestBuiltinFunctionType2(t *testing.T) {
	bs1 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	bs2 := bs1[10:]
	var x byte
	x = 1 << copy(bs2, bs1) / 128
	fmt.Println(x) // 0
}

func TestShiftReturnType(t *testing.T) {
	const c9 = 9
	v9 := c9

	var x, y byte
	x = 1 << c9 / 128
	y = 1 << v9 / 128
	fmt.Println(x, y) // 4 0
}

func TestShiftReturnType2(t *testing.T) {
	const c9 = 9
	v9 := c9
	fmt.Println(byte(1<<c9/128), byte(1<<v9/128)) // 4 0
}

func TestShiftReturnType3(t *testing.T) {
	const c9 = 9
	var v9 = 9
	var a byte = 1 << c9 >> c9
	var b byte = 1 << v9 >> c9
	var c byte = 1 << v9 >> v9
	fmt.Println(a, b, c)
}

// a. 0 0 0
// b. 1 0 0
// c. 1 1 0
// d. 1 1 1
