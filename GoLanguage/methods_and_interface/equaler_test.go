package __

import (
	"log"
	"testing"
)

// Equaler is an attempt to create an interface that supports comparison (equality)
// TODO: use generics now that they are available
type Equaler interface {
	Equals(interface{}) bool
}

// Integer is an int type that supports Equaler interface
type Integer int

//// Equals version that works for Integer and int
//func (v1 Integer) Equals(v2 interface{}) bool {
//	a2, ok := v2.(Integer)
//	if !ok {
//		b2, ok := v2.(int)
//		if !ok {
//			return false
//		}
//		return int(v1) == b2
//	}
//	return v1 == a2
//}

// Equals version that is more general (easily extended)
func (v1 Integer) Equals(v2 interface{}) bool {
	switch v2 := v2.(type) {
	case Integer:
		return v1 == v2
	case int:
		return int(v1) == v2
		// extend here for other types
		//case string: // we don't really want this
		//	i, err := strconv.Atoi(v2)
		//	if err != nil {
		//		return false
		//	}
		//	return int(v1) == i
	}
	return false
}

func TestEqualer(t *testing.T) {
	var e Equaler
	e = Integer(1)
	log.Println(e.Equals(Integer(1))) // true
	log.Println(e.Equals(Integer(2))) // false
	log.Println(e.Equals(1))          // true
	log.Println(e.Equals(2))          // false
	//log.Printf("%t\n", 1.Equals(Integer(2)))           // int does not support Equals
}
