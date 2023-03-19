package __

import (
	"log"
	"testing"
)

// See also GoLanguage/math/literals_test.go

func TestNestedStruct(t *testing.T) {
	type axx struct{ int }
	type bxx struct {
		axx
		string
	}

	a := axx{10}
	log.Println(a)          // {10}
	log.Println(a.int)      // 10
	b := bxx{axx{5}, "str"} //
	log.Println(b)          // {{5} str}
	log.Println(b.axx.int)  // 5
}

func TestStructLiteral(t *testing.T) {
	type a struct{ b, c int } //
	d := a{2, 3}              // by order
	log.Println(d)            // {2 3}
	e := a{c: 4}              // by field name (safer)
	log.Println(e)            // {0 4}
	f := &a{5, 6}             // => new
	log.Println(f)            // &{5 6}
}

func TestArray(t *testing.T) {
	type h [4]int
	i := h{1, 2, 3}          // by order
	log.Println(i)           // [1 2 3 0]
	j := h{2: 22, 0: 33}     // by index (cool)
	log.Println(j)           // [33 0 22 0]
	k := &h{5, 6, 7, 8}      // => new
	log.Println(k)           // &[5 6 7 8]
	l := &[4]int{5, 6, 7, 8} //
	log.Println(l)           // &[5 6 7 8]
}

func TestSliceLiteral(t *testing.T) {
	n1 := []int{}
	log.Println(n1)               // []
	n2 := []int{1, 2, 3, 1, 4, 5} // specify elements in order
	log.Println(n2)               // [1 2 3 1 4 5]
	n3 := []int{7: 42, 2: 22}     // specify elements by index (cf map literal)
	log.Println(n3)               // [0 0 22 0 0 0 0 42]
	p1 := &[]int{42}              // Note that this calls new to get a ptr to slice, whereas make([]int, 1) returns a slice
	log.Println(p1)               // &[42]
	p2 := &[]int{4, 2}            //
	log.Println(p2)               // &[4 2]
	p3 := make([]int, 3)          //
	log.Println(p3)               // [0 0 0]
}

func TestMapLiteral(t *testing.T) {
	type q map[int]string
	r := q{2: "22", 0: "33"}
	log.Println(r)              // map[2:22 0:33]
	s := &q{}                   // Note that this calls new() to get a ptr to map, whereas make() returns a map (ref type)
	log.Println(s)              // &map[]
	u := &q{888: "a", 999: "b"} //
	log.Println(u)              // &map[888:a 999:b]
	v := make(q, 5)             //
	log.Println(v)              // map[]
}

func TestLiteralInt(t *testing.T) {
	type aa int
	bb := aa(2)                   // aa "literal"
	log.Printf("%T %v\n", bb, bb) // aa 2
	//cc := &aa(3)                  // apparently using & only works for composite types
	cc := new(aa)
	*cc = aa(3)
	log.Println(cc, *cc) // 0xc042008f90 3
}
func TestLiteralInterface(t *testing.T) {
	// TODO
}
