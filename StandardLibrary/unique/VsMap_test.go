package __

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	uHandle[T comparable] struct {
		h *T
	}
	uMap[T comparable] map[T]uHandle[T]
)

func (h uHandle[T]) Value() T {
	return *h.h
}

var allMaps = make(map[reflect.Type]any)

func uMake[T comparable](s T) uHandle[T] {
	var umap uMap[T]

	// First get the map for type T
	t := reflect.TypeOf(s)
	theMap, ok := allMaps[t]
	if !ok {
		// Create and add it
		umap = make(uMap[T])
		allMaps[t] = umap
	} else {
		// Found but must be converted to our specific map type
		umap = theMap.(uMap[T])
	}

	if r, found := umap[s]; !found {
		//s = strings.Clone(s)
		r = uHandle[T]{h: &s}
		umap[s] = r
		return r
	} else {
		return r
	}
}

func TestUniqueMapString(t *testing.T) {
	h1 := uMake("abc")
	h2 := uMake("ab" + "c")
	h3 := uMake("def")

	fmt.Println(h1, h1.Value())
	fmt.Println(h2, h2.Value())
	fmt.Println(h3, h3.Value())
}

func TestUniqueMap2(t *testing.T) {
	type MyString string

	var ms MyString = "abc"
	h1 := uMake("abc")
	h2 := uMake(MyString("abc"))
	h3 := uMake[MyString]("abc")
	h4 := uMake(ms)

	fmt.Println(h1, h1.Value())
	fmt.Println(h2, h2.Value())
	fmt.Println(h3, h3.Value())
	fmt.Println(h4, h4.Value())
}

/*
func uMake[T comparable](s T) uHandle[T] {
	t := reflect.TypeOf(s)
	theMap, ok := allMaps[t]
	if !ok {
		theMap = make(uMap[T])
		allMaps[t] = theMap
	}

	umap, ok := theMap.(uMap[T])
	if !ok {
		panic("only doing strings for now")
	}
	if r, found := umap[s]; !found {
		//s = strings.Clone(s)
		r = uHandle[T]{h: &s}
		umap[s] = r
		return r
	} else {
		return r
	}
}
*/
