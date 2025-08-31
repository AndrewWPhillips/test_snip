package utils

import (
	"log"
	"testing"
	"time"
)

// These are tests of functions which can be made generic when we have a Go 1.18 RC

// Min return the minimum of 1 or more ints - with generics extend to any comparable type (byte, float64, string)
func Min(first int, rest ...int) (retval int) {
	retval = first
	for _, v := range rest {
		if v < retval {
			retval = v
		}
	}
	return
}

func TestMin(t *testing.T) {
	log.Println(Min(77, 24, 42, 99, 73, 101, 24))
}

// Stringify converts a slice of time.Times to a slice of string
// With generics it could take a type parameter (any type that implements Stringer)
func Stringify(s []time.Time) (retval []string) {
	for _, v := range s {
		retval = append(retval, v.String())
	}
	return
}

// This uses old go2go generics with contracts
//type Stringer struct {}
//func (s Stringer) String() string { return "" }
//
//func Stringify(type Stringer)(s []Stringer) (ret []string) {
//	for _, v := range s {
//		ret = append(ret, v.String())
//	}
//	return ret
//}

// This uses final generics proposal (with constraints and square brackets)
//func Stringify[T interface{ String() string }](s []T) []string {
//	r := make([]string, len(s))
//	for _, v := range s {
//		r = append(r, v.String())
//	}
//	return r
//}

func TestStringify(t *testing.T) {
	vals := []time.Time{time.Now(), time.Now().Add(25 * time.Hour)}
	strs := Stringify(vals)
	log.Println(strs)
}
