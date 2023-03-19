package __

import (
	"fmt"
	"log"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []string{"abc", "def", "ghi"}
	for _, v := range s {
		v = v + "??" // does not modify slice
	}
	log.Println(s) // [abc def ghi]

	ss := make([][]string, 3)
	ss[0] = []string{"abc"}
	ss[1] = []string{"def"}
	ss[2] = []string{"ghi"}
	for _, v := range ss {
		v[0] = v[0] + "??" // v[0] points to same thing as ss[idx][0]
	}
	log.Println(ss) // [[abc??] [def??] [ghi??]]
}

// Tests modifying an array in a loop and using the array values
func TestModification(t *testing.T) {
	var a [10]int
	//a := make([]int, 10)

	a[0] = 1
	for i, v := range a {
		next := i + 1
		if next < len(a) {
			a[next] = v * next
		}
	}
	log.Println(a)
	// array: 1 1 0 0 0 0 0 0 0 0 which shows that a copy of a is made to range over
	// slice: 1 1 2 6 24 120 720 5040 40320 362880
}

func TestRangeLoopVarsDeclared(t *testing.T) {
	var i int
	var v int
	for i, v = range []int{1, 2, 3, 4, 5} {
		if i == 2 {
			break
		}
	}
	log.Println(i, v)
}

func TestWrongRangeInt(t *testing.T) {
	a := []int{1, 2, 3}
	sum := 0
	for v := range a {
		//for _, v := range a { // this is what was intended
		sum += v
		// should get: invalid operation: sum += v (mismatched types int and loop index)
	}
	log.Println(sum) // 3 (not 6) since we are summing the indices not the values
}

func TestWrongRangeString(t *testing.T) {
	a := []string{"10", "20", "30"}
	sum := ""
	//for v := range a { // invalid operation: sum += v (mismatched types string and int)
	for _, v := range a {
		sum += v
	}
	log.Println(sum) // "102030"
}

func TestModifyIndex(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	i := 42
	pi := &i
	for i = range a {
		log.Println(*pi)
		*pi = 0
		a[i] = 4
	}
	log.Println(a, i)
}

func TestString(t *testing.T) {
	for i, c := range "a¿b" {
		fmt.Println(i, c, string(c))
	}
}
