package __

import (
	"reflect"
	"runtime"
	"testing"
	"time"
	"unique"
)

// TestUniqueHoldStringField is like TestUniqueHoldSubString but checks that a substring as
// a struct field does not prevent the larger string from being GCed
func TestUniqueHoldStringField(t *testing.T) {
	a := "abcdefhijk"
	unique.Make(struct{ string }{a[1:2]})

	runtime.AddCleanup(&a, func(i int) { println("GC of a") }, 0)
	a = "" // ensure the original string is no longer referenced
	println("before GC")
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
}

// TestUniqueHoldSubNotString is the same as TestUniqueHoldSubString but uses
// a new string type - this test that cloning works based on the underlying type
func TestUniqueHoldSubNotString(t *testing.T) {
	type MyString string
	a := "abcdefhijk"
	unique.Make(MyString(a[1:2]))

	// unique uses the underlying type to find the string fields.
	// That is it calls reflect.Type.Kind() (see src/unique/clone.go:46,60,77)
	println(reflect.TypeOf(a).Kind(), reflect.TypeOf(MyString(a)).Kind())

	runtime.AddCleanup(&a, func(i int) { println("GC of a") }, 0)
	a = "" // ensure the original string is no longer referenced
	println("before GC")
	runtime.GC()
	time.Sleep(time.Microsecond) // allow for GC to finish
	println("after GC")
}

func TestUniqueHoldStringPtrField(t *testing.T) {
	a :=
		"iuirueytiureyitrueyituryeitureytirueytoreiutorieyugtoreiutorieutroeiutroeiutreoiutreoiutreoiture" +
			"iuirueytiureyitrueyituryeitureytirueytoreiutorieyugtoreiutorieutroeiutroeiutreoiutreoiutreoiture" +
			"ytreoutreoutroeiutroeiutreoiutreoiutreoiutreoiutreoiutreoiutreoiutreoitureoitureoitureoiutreoiutreeoi"
	s1 := a[1:2]
	h1 := unique.Make(&s1)

	runtime.AddCleanup(&a, func(i int) { println("GC of a") }, 0)
	a = "" // ensure the original string is no longer referenced
	println("before GC")
	runtime.GC()
	time.Sleep(time.Second) // allow for GC to finish
	println("after GC")
	runtime.KeepAlive(h1)
}
