package __

// use go generate to run stringer to generates a "string" method on the MyEnum type (see myenum_string.go)
//go:generate stringer -type=MyEnum

// Note that I could not put this is myenum_test.go as //go:generate did not seem to work from there

type MyEnum byte

const (
	Zero MyEnum = iota
	One
	Two
	Three

	// MyEnumLast is the value one past the last valid enum value
	MyEnumLast
	// MyEnumFirst is the first valid enum value
	MyEnumFirst = Zero
)

const Uno MyEnum = 1

// MyEnumRange can be used to iterate all the enum values using a for loop like this:
//   for i := range MyEnumRange {
//      v := MyEnum(i)
//      ...
var MyEnumRange [MyEnumLast]struct{}
