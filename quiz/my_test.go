package __

import (
	"fmt"
	"testing"
)

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
