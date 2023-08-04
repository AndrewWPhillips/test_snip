package __

import (
	"bytes"
	"log"
	"math"
	"testing"
)

// TestArrayKey tries to (unsuccessfully) confuse a map by modifying a key of array type
func TestArrayKey(t *testing.T) {
	m := map[[2]int]string{
		[2]int{0, 0}: "0",
		[2]int{1, 2}: "1",
	}
	a := [2]int{1, 2}
	pa := &a
	m[a] = "2"
	log.Println(m)
	*pa = [2]int{0, 1} // modifies a, but not the key in m
	log.Println(m)
}

func TestArrayKey2(t *testing.T) {
	type a10_t [10]int
	var a1, a2 a10_t

	a1[5] = 55
	a2 = a1
	log.Println(a1, a2, a1 == a2) // T
	a2[3] = 33
	log.Println(a1, a2, a1 == a2) // F

	m := map[a10_t]string{a1: "aaaaa", a2: "bbbbb"}

	log.Println(m[a1], m[a2]) // aaaaa bbbbb
	a2 = a1
	log.Println(m[a1], m[a2]) // aaaaa aaaaa
	// Note: book "Way to Go" says you can't use an array as a map key but you clearly can (since arrays can be compared)
}

// TestNanKey shows the strange things that happen when you use a NaN as a map[float64]T key value
func TestNanKey(t *testing.T) {
	m := map[float64]int{
		math.NaN(): 1,
		math.Pi:    2,
	}
	log.Println(m[math.NaN()], m[math.Pi])
}

func TestNanKey2(t *testing.T) {
	m := make(map[float64]int)

	for i := 1; i < 10; i++ {
		m[math.NaN()] = i
	}
	for k := range m {
		log.Println(m[k])
	}
}

func TestStructKey(t *testing.T) {
	type aa struct {
		bb int16
		cc int32
	}
	key := &aa{}
	ma := make(map[*aa]string) // use ptr to struct as key

	ma[key] = "yyy"
	log.Println(ma[key]) // "yyy"
	key.bb = 2
	log.Println(ma[key]) // "yyy"

	mb := map[aa]string{*key: "zzz"} // use struct as key

	log.Println(mb[*key]) // "zzz"
	key.bb = 3
	log.Println(mb[*key]) // ""
}

func TestStructKey2(t *testing.T) {
	var a, b struct{}
	log.Println(a == b)                        // true - all empty structs are equal
	m1 := map[struct{}]bool{a: false, b: true} // a == b so only one element is added
	log.Println(m1)                            // map[{}:true]

	m2 := map[struct{ i int }]string{struct{ i int }{43}: "aaaa", struct{ i int }{42}: "bbbb"}
	log.Println(m2) // map[{42}:bbbb {43}:aaaa]

	type s struct{ i int }
	m3 := map[s]string{s{1}: "aaaa", s{2}: "bbbb"}
	log.Println(m3) // map[{1}:aaaa {2}:bbbb]
}

func TestBinKey(t *testing.T) {
	mm := make(map[[3]byte]int)

	a := [3]byte{1, 2, 3}
	mm[a] = 42
	aa := [3]byte{1, 2, 3}
	mm[aa] = 77
	b := [3]byte{1, 2, 5}
	mm[b] = 43
	log.Println(mm)    // map[[1,2,3]:77, [1,2,5]:43]
	log.Println(mm[a]) // 77 - hence mm[a] is the same elt as mm[aa] so hash is on the contents of the array key

	// mess with the map key value
	a[2] = 5
	log.Println(mm) // does not cause a problem hence map must have a copy of the key
}

func TestPointerKey(t *testing.T) {
	i, j, k := 42, 43, 42
	m := map[*int]string{&i: "iii", &j: "jjj", &k: "kkk"}
	log.Println(m) // map[0xc042008e48:kkk 0xc042008e38:iii 0xc042008e40:jjj]
}

func TestInterfaceKey(t *testing.T) {
	// Can't use: SMF (slice, map, func) as key (even if interface)
	// Can use: chan, interface, pointer, numeric, etc (comparable types)
	// Maybe: arrays and structs as long as they only contain comparable types
	var a, b, c interface{}
	a = 1
	b = "2"
	c = 3.3
	m := map[interface{}]int{a: 1, b: 2, c: 3}
	log.Printf("%v\n", m) // map[3.3:3 1:1 2:2]
}

func TestSMFInterfaceKey(t *testing.T) {
	m := map[interface{}]int{interface{}([]int{}): 1} // panic: runtime error: hash of unhashable type []int
	_ = m
}

func TestModifyKey(t *testing.T) {
	mm := map[string]int{"1,2": 1, "3": 3}
	log.Println(mm) // map[1,2:1 3:3]

	for k, v := range mm {
		if k == "3" || v == 3 {
			k = "42"
		}
	}
	log.Println(mm) // map[1,2:1 3:3]
}

func TestQuizQ18(t *testing.T) {
	s := bytes.NewBufferString("string")
	a := s.String()
	b := s.Bytes()
	m := make(map[interface{}]int)
	m[a] = 1
	m[b] = 2
	log.Println(m[a])
}
