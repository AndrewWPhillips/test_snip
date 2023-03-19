package __

import (
	"reflect"
	"testing"
)

func TestFuncEquality(t *testing.T) {
	a := TestFuncEquality
	b := TestFuncEquality
	t.Log(reflect.DeepEqual(a, b))
	a = nil
	b = nil
	t.Log(reflect.DeepEqual(a, b))
}
