package __

import (
	"log"
	"testing"
)

// Go "enums" are criticised for these problems:
//  1. do not enforce the range of values (type safety)
//  2. do not have a way to iterate over valid values,
//  3. have no auto string conversion
//  4. no distinct identity
// These issues can be handled by carefully crafting and using a type like MyEnum (see
// myenum.go) and using go generate (see generated file myenum_string.go)

// TestIteration shows how to iterate over all the enum values (prob 2) and print as string (prob 3)
func TestValues(t *testing.T) {
	//for v := MyEnumFirst; v < MyEnumLast; v++ {
	//	log.Println(v)
	//}

	for i := range MyEnumRange {
		v := MyEnum(i)
		log.Println(i, v)
	}
}

// TestIllegalValues shows how easily an "illegal" value can be assigned to an "enum" variable
func TestIllegalValues(t *testing.T) {
	var e MyEnum
	e = 42 // don't assign an integer literal to an "enum" (prob 1)
	log.Println(e)
}

// TestDupeEnumValue shows how you can add two enum "values" with the same integer value
func TestDupeEnumValue(t *testing.T) {
	// This enum "value" is not added by go generate - I think *_test.go files are ingored
	const Dupe MyEnum = 2 // same value as "Two"

	log.Println(Uno)         // One
	log.Println(Uno == One)  // true
	log.Println(Dupe == Two) // true
}
