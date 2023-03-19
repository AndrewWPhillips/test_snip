package __

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestArrayVsSlice(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	log.Println(reflect.TypeOf(a))           // array of 5 ints
	log.Println(reflect.TypeOf(b))           // slice of ints
	log.Printf("a: %d %d\n", len(a), cap(a)) // 5 5
	log.Printf("b: %d %d\n", len(b), cap(b)) // 5 5

	b = b[1:3]
	log.Printf("b: %d %d\n", len(b), cap(b)) // 2 4
}

// TestArrayReturn shows that returning an array results in value that's not addressable
func TestArrayReturn(t *testing.T) {
	// Here is a function that returns an array
	f := func() [4]int { return [4]int{1, 2, 3, 4} }

	// Can't slice an array that is returned
	//fmt.Println(f()[:2]) // invalid operation f()[:2] (slice of unaddressable value)

	// Assign to a temporary to allow slicing
	tmp := f()
	fmt.Println(tmp[:2]) // OK, prints: [1 2]
}

func TestArrayLiteral(t *testing.T) {
	// Can't slice an array literal either
	//fmt.Println([4]int{1, 2, 3, 4}[1:2]) // invalid operation [4]int literal[1:2] (slice of unaddressable value)

	tmp := [4]int{1, 2, 3, 4}
	fmt.Println(tmp[1:2])
}

func TestArrayAssign(t *testing.T) {
	var a1 = [3]int{42, 43, 44}
	a2 := a1
	a1[1] = 99
	log.Printf("%v   %v\n", a1, a2)
	// Different values => assigning an array makes a copy
}

func Test2dArray(t *testing.T) {
	var a [2][3]int

	//for i := 0; i < 2; i++ {
	//	for j := 0; j < 3; j++ {
	//		a[i][j] = (i + 1) * (j + 1)
	//	}
	//}
	for i := range a {
		for j := range a[i] {
			a[i][j] = (i + 1) * (j + 1)
		}
	}
	log.Println(a) // [[1 2 3] [2 4 6]]

	a = [2][3]int{{2, 4, 6}, {3, 6, 9}}
	log.Printf("%v\n", a) // [[2 4 6] [3 6 9]
}
