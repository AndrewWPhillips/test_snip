package __

import (
	"fmt"
	"log"
	"testing"
)

func TestAddress(t *testing.T) {
	i := int16(42)
	j := int64(1e9)
	k := i
	s1 := "abc"
	s2 := s1
	m := make(map[string]int)
	n := map[string]int{"abc": 1, "def": 2}
	p := m
	fmt.Printf("i %p\n", &i)
	fmt.Printf("j %p\n", &j)
	fmt.Printf("k %p\n", &k)
	fmt.Printf("s1 %p\n", &s1)
	fmt.Printf("s2 %p\n", &s2)
	fmt.Printf("m %p\n", &m)
	fmt.Printf("n %p\n", &n)
	fmt.Printf("p %p\n", &p)
}

func TestAddressofLiteral(t *testing.T) {
	//i := &42
	a := &[1]int{1}                // array
	s := &[]int{1, 2}              // slice
	m := &map[string]int{"": 0}    // map
	u := &struct{ x, y int }{1, 2} // struct
	v := &struct{}{}
	//z := &interface{}{}
	fmt.Printf("a %T %v\n", a, a)
	fmt.Printf("s %T %v\n", s, s)
	fmt.Printf("m %T %v\n", m, m)
	fmt.Printf("u %T %v\n", u, u)
	fmt.Printf("v %T %v\n", v, v)
}

// TestSliceP shows that %p for a slice is the address of the data (the ptr in the slice header)
func TestSliceP(t *testing.T) {
	si := []int{1, 2, 3}
	log.Printf("%p %p %p %p\n", si, &si[0], &si[1], &si) // si and &si[0] give the same printed address
}

// TestMapP shows that %p fr a map gives the address of the data not the map itself
func TestMapP(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	//log.Printf("%p %p\n", m, &m["one"]) // invalid operation: cannot take address...
	m2 := m
	log.Printf("%p %p\n", m, m2)
	log.Printf("%p %p\n", &m, &m2)
}

func TestStructP(t *testing.T) {
	u := struct{ x, y int }{1, 2}
	//log.Printf("%p\n", u) // can't use %p with struct (only ptr, map, slice, chan, func)
	log.Printf("%p %p\n", &u, &u.x)
}

func TestChanP(t *testing.T) {
	c := make(chan int)
	log.Printf("%p %p\n", &c, c)
}

func TestFuncP(t *testing.T) {
	f := func() {}
	log.Printf("%p %p\n", &f, f)
}

func TestInterfaceP(t *testing.T) {
	z := any(2)
	log.Printf("%p %p\n", &z, z) // 0xc000026320 %!p(int=2)
	z2 := any(make(chan int))
	log.Printf("%p %p\n", &z2, z2) // 0xc000026330 0xc00005c230
}
