package maps

import (
	"log"
	"testing"
)

func TestReturn2FromMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	mm := func(s string) (int, bool) {
		//return m[s]
		a, b := m[s]
		return a, b
	}

	log.Println(mm("one"))
	log.Println(mm("two"))
	log.Println(mm("tree"))

}
