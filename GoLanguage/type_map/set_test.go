package __

import (
	"log"
	"testing"
)

type (
	PlaceHolder struct{}
	StringSet   map[string]PlaceHolder
)

func Intersect(set ...StringSet) StringSet {
	retval := make(StringSet)
	if len(set) == 0 {
		return retval
	}
outer:
	for elt := range set[0] {
		for _, other := range set[1:] {
			if _, ok := other[elt]; !ok {
				continue outer
			}
		}
		retval[elt] = PlaceHolder{}
	}
	return retval
}

func TestSetIntersect(t *testing.T) {
	setA := StringSet{"a": {}, "b": {}, "c": {}}
	setB := StringSet{"c": {}, "d": {}, "e": {}, "b": {}}
	setC := StringSet{"e": {}, "a": {}, "c": {}, "b": {}}

	log.Println(Intersect(setA, setB, setC))

}
