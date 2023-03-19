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
	fmt.Printf("IntSlice %p\n", &m)
	fmt.Printf("n %p\n", &n)
	fmt.Printf("p %p\n", &p)
}

// TestSlice checks if %p for a slice is the address of the slice "header" (addr+len+cap) or the data
func TestSlice(t *testing.T) {
	si := []int{1, 2, 3}
	log.Printf("%p %p %p\n", si, &si[0], &si[1]) // si and &si[0] give the same printed address
}

func TestMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	//log.Printf("%p %p\n", m, &m["one"]) // invalid operation: cannot take address...
	m2 := m
	log.Printf("%p %p\n", m, m2)
}
