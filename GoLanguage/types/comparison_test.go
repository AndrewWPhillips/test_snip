package __

import (
	"fmt"
	"log"
	"testing"
)

// TestPIC compares pointers, interfaces and chans ("ref" types that can be compared to each other)
func TestPIC(t *testing.T) {
	i, j := 42, 42
	pi, pj := &i, &j
	log.Println(pi == pi, pi == pj) // T F

	log.Println(interface{}(i) == interface{}(j)) // true - does not compare ptrs but looks at underlying types/values

	ci, cj := make(chan int), make(chan int)
	log.Println(ci == ci, ci == cj) // T F
}

func TestAnonStruct(t *testing.T) {
	a := struct{ v int }{v: 1}
	b := struct{ v int }{v: 2}
	c := struct{ v int }{v: 2}
	log.Println(a == b, b == c) // F T
}

func TestInterface(t *testing.T) {
	var x, y interface{}
	log.Println(x == y) // true - both nil

	x = struct{}{}
	y = struct{}{}
	log.Println(x == y) // true - same type and (comparable) value

	x = struct{ s []int }{}
	y = struct{}{}
	log.Println(x == y) // false - diff type

	x = struct{ s []int }{}
	y = struct{ s []int }{}
	log.Println(x == y) // panic - same (noncomparable) type
}

func TestInterfaceEqInt(t *testing.T) {
	var x interface{}
	x = int32(0)
	if x == 0 {
		log.Println("this doesn't happen!")
	}
	if x.(int32) == 0 {
		log.Println("this does happen")
	}
	if x.(int32) == 42 {
		log.Println("this doesn't happen")
	}
	if x.(int) == 42 {
		log.Println("this panics")
	}
}

func TestMapPanic(t *testing.T) {
	// If a struct contains non-comparable type then we can't compare the struct (or use as map key)
	// A non-comparable type = slice/map/func (or struct/array/interface containing non-comparable type)
	type test struct {
		i int
		//x []int // slice -> we can't use this struct as map key
		//x [0][]int // array of slice
		//x struct{ f func() } // struct with func
		y *int // PIC - can compare pointer/interface/chan
		z interface{}
	}
	a := test{i: 1}
	b := test{i: 2}
	m := make(map[interface{}]int)
	m[a] = 1 // panic unless we remove the x (slice) from the struct
	m[b] = 2
	log.Println(m[a])
}

func TestNil(t *testing.T) {
	var f, g *int
	var h interface{}
	var i interface{} = f

	log.Println(g == f)   // true
	log.Println(h == f)   // false - h is not of *int type
	log.Println(i == f)   // true - both of same type (*int) with nil value
	log.Println(g == nil) // true
	log.Println(h == nil) // true
	log.Println(i == nil) // false - i is of *int type with nil value
	// An interface is == nil if it has a nil value and no type BUT it can be equal to a nil
	// value of another type if it has a nil value and the same type
}

var emptyA struct{}

func TestEmptyStruct(t *testing.T) {
	var emptyB struct{}
	p1, p2 := &emptyA, &emptyB
	fmt.Printf("&a: %p, &b: %p \n", &emptyA, &emptyB)
	fmt.Printf("&a == &b: %v \n", &emptyA == &emptyB)
	fmt.Printf("p1 == p2: %v \n", p1 == p2)
	fmt.Printf("%p %p\n", p1, p2)
	emptyComp(p1, p2)
	emptyComp2(p1, p2)
}

func emptyComp(p1, p2 *struct{}) {
	fmt.Printf("p1 == p2: %v \n", p1 == p2)
	fmt.Printf("%p %p\n", p1, p2)
}

func emptyComp2(p1, p2 any) {
	fmt.Printf("p1 == p2: %v \n", p1 == p2)
	fmt.Printf("%p %p\n", p1, p2)
}
