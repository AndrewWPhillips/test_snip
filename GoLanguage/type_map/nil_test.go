package __

import (
	"log"
	"testing"
)

func TestNestedNilMaps(t *testing.T) {
	var m map[int]map[string]map[struct{}]bool

	if _, ok := m[1][""][struct{}{}]; ok { // false: m is nil
		delete(m[1][""], struct{}{})
	}
	if _, ok := m[1][""]; ok {
		delete(m[1], "")
	}
	if _, ok := m[1]; ok {
		delete(m, 1)
	}

	m = make(map[int]map[string]map[struct{}]bool)
	if _, ok := m[1][""][struct{}{}]; ok { // false: m[1] is nil
		delete(m[1][""], struct{}{})
	}
	if _, ok := m[1][""]; ok {
		delete(m[1], "")
	}
	if _, ok := m[1]; ok { // false m[1] not found
		delete(m, 1)
	}

	m = make(map[int]map[string]map[struct{}]bool)
	m[1] = make(map[string]map[struct{}]bool)
	m[1][""] = make(map[struct{}]bool)
	m[1][""][struct{}{}] = true
	if _, ok := m[1][""][struct{}{}]; ok { // true
		delete(m[1][""], struct{}{})
	}
}

func TestRead(t *testing.T) {
	var m map[string]int
	log.Println(m[""]) // read from nil map returns zero value

	m = make(map[string]int)
	log.Println(m[""]) // read non-existent element gives zero value

	m[""] = 42
	log.Println(m[""])
}

func TestWrite(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	log.Println(m) // read non-existent element gives zero value

	//m = nil
	//m["a"] = 1 // assignment to nil map entry
}

func TestLen(t *testing.T) {
	p := new(struct{ allow []string })
	log.Println(len(p.allow)) // 0 since p.allow == nil

	p.allow = make([]string, 3)
	log.Println(len(p.allow)) // 3

	p.allow = make([]string, 0, 3)
	log.Println(len(p.allow)) // 0

	//p = nil
	//log.Println(len(p.allow)) // panic

	var m map[string]int
	log.Println(len(m))

	var ch chan int
	log.Println(len(ch))
}

func TestAll(t *testing.T) {
	// You can do almost anything with a nil map but assign to an element
	// - delete of an element does nothing
	// - access an element returns default (zero) value for element
	var mm map[int]string
	delete(mm, 1)          // OK: does nothing
	for k, v := range mm { // OK: loop executed zero times
		log.Println(k, v)
	}
	log.Println(len(mm)) // OK: 0
	log.Println(mm[1])   // OK: empty string *** inconsistent with slice/array
	v, got := mm[1]      // OK
	log.Println(v, got)  // empty string, false
	log.Println(mm)      // OK: map[]
	mm[1] = "abc"        // panic!!!!!
}
