package __

import (
	"fmt"
	"math"
	"testing"
)

// TestNanMap
// What does this print?
// a. 0
// b. 1
// c. 2
func TestNanMap(t *testing.T) {
	m := make(map[float64]int)
	m[math.NaN()] = 1
	m[math.NaN()] = 2
	delete(m, math.NaN())
	fmt.Println(len(m))
}

// c because math.Nan() != math.NaN()

func TestForEndValue(t *testing.T) {
	var i, I int
	for i = 0; i < 10; i++ {
	}
	for I = range [10]int{} {
	}
	fmt.Println(i, I)
}

// TestDeferArgEvalTime test when args to a deferred function are evaluated
// What does this print?
// a. 111
// b. 121
// c. 122
// d. 211
// e. 222
func TestDeferArgEvalTime(t *testing.T) {
	i := 1
	defer fmt.Print(i)
	defer func() {
		fmt.Print(i)
	}()
	defer func(i int) {
		fmt.Print(i)
	}(i)
	i++
}

// answer: b (1st/3rd: i is evaluated when defer is called, for 2nd: i is captured)

func f() error {
	return nil // errors.New("error")
}

// What does this print?
// a. <nil>
// b. *** not nil ***
// c. does not compile
// d. func() error
func TestNilFuncInInterface(t *testing.T) {
	if e := f; e == nil {
		fmt.Printf("%T", e)
	} else {
		fmt.Print("*** not nil ***")
	}
}

// answer: b

// What does this print?
// a. 0 0 0
// b. 1 0 0
// c. 1 1 0
// d. 1 1 1
func TestShifty(t *testing.T) {
	const c9 = 9
	var v9 = 9
	var a byte = 1 << c9 >> c9
	var b byte = 1 << v9 >> c9
	var c byte = 1 << v9 >> v9
	fmt.Println(a, b, c)
}

// b. because 1 << v9 is treated as byte (because the expr. is assigned to a byte and the rest of the expr. only has constants)
