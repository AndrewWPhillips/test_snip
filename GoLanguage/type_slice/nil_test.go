package __

import (
	"log"
	"testing"
)

func TestNilSlice(t *testing.T) {
	// You can do almost anything with a nil slice
	// - can't index into it but can't do that on zero length slice anyway
	// - copy effectively does nothing if either source or dest is nil
	// - append may allocate memory
	var ss []int
	for i := range ss { // OK: loop executed zero times
		log.Println(i)
	}
	for i, v := range ss { // OK: loop executed zero times
		log.Println(i, v)
	}
	log.Println(len(ss), cap(ss))      // OK: 0  0
	log.Println(copy(ss, ss))          // OK: 0
	log.Println(copy([]int{1, 2}, ss)) // OK: 0
	log.Println(copy(ss, []int{1, 2})) // OK: 0
	log.Println(ss)                    // OK: []
	log.Println(ss[:])                 // OK: []
	log.Println(ss[:0])                // OK: []
	log.Println(append(ss, 1, 2))      // OK: [1 2]

	//log.Println(ss[:1]) // panic: runtime error: slice bounds out of range [:1] with capacity 0
	//log.Println(ss[0]) // panic  *** inconsistent with map
	//ss[0] = 42 // panic!!
}

type MyError struct{}

func (MyError) Error() string { return "ERROR" }

func stub() *MyError { return nil }

func TestNilError(t *testing.T) {
	var err error = stub()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("OK")
}
