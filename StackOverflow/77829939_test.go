package __

import (
	"fmt"
	"reflect"
	"testing"
)

func creator(v string) func() string {
	return func() string {
		return v
	}
}

func getters(runID string) []func() string {
	return []func() string{
		creator(runID),
	}
}

func TestPtrs(t *testing.T) {
	var allVals []func() string
	for i := 0; i < 20; i++ {
		vals := getters(fmt.Sprintf("run-%d", i))
		for j := range vals {
			fmt.Printf("[%d] Valueof %v pointer: %x\n", i, reflect.ValueOf(vals[j]), reflect.ValueOf(vals[j]).Pointer())
		}
		allVals = append(allVals, vals...)
	}

	for i := range allVals {
		fmt.Printf("[%d] %s\n", i, allVals[i]()) // Works as expected
	}
}

func TestPtrs1(t *testing.T) {
	for i := 0; i < 3; i++ {
		vals := getters(fmt.Sprintf("run-%d", i))
		value := reflect.ValueOf(vals[0])
		fmt.Println(value)
		fmt.Println(value.Pointer())
	}
}

// Same address for each value
func TestPtrs2(t *testing.T) {
	var val func() string
	var value reflect.Value
	for i := 0; i < 2; i++ {
		val = creator("")
		value = reflect.ValueOf(val)
		fmt.Println(value)
		fmt.Println(value.Pointer())
	}
}

func get() func() int {
	return func() int {
		return 42
	}
}

func TestPtrs3(t *testing.T) {
	var val func() int
	var value reflect.Value
	for i := 0; i < 2; i++ {
		val = get()
		value = reflect.ValueOf(val)
		fmt.Println(value)
		fmt.Println(value.Pointer())
	}
}

// 15876032
// 15876032
func TestPtrs4(t *testing.T) {
	for i := 0; i < 2; i++ {
		fmt.Println(reflect.ValueOf(creator("")).Pointer())
	}
}

// ??? WTF

func TestPtrs5(t *testing.T) {
	for i := 0; i < 2; i++ {
		fmt.Println(reflect.ValueOf(creator("")))
		fmt.Println(reflect.ValueOf(creator("")).Pointer())
	}
	fmt.Println(reflect.ValueOf(creator("")))
	fmt.Println(reflect.ValueOf(creator("")).Pointer())
	fmt.Println(reflect.ValueOf(creator("")))
	fmt.Println(reflect.ValueOf(creator("")).Pointer())
}

// TestPtrs6 unrolls the loop of TestPtrs4 but prints 2 different values
// but if get() is not inlined then it prints the same value twice
func TestPtrs6(t *testing.T) {
	fmt.Println(reflect.ValueOf(creator("")).Pointer())
	fmt.Println(reflect.ValueOf(creator("")).Pointer())
}

// Same as TestPtrs3 except the loop was unrolled but gives:
// 0x1013d40
// 16858432
// 0x1013d60
// 16858464
func TestPtrs7(t *testing.T) {
	var val func() int
	var value reflect.Value
	val = get()
	value = reflect.ValueOf(val)
	fmt.Println(value)
	fmt.Println(value.Pointer())
	val = get()
	value = reflect.ValueOf(val)
	fmt.Println(value)
	fmt.Println(value.Pointer())
}

// Diff address for each value
func TestPtrs8(t *testing.T) {
	var val func() string
	var value reflect.Value
	for i := 0; i < 2; i++ {
		if i == 0 {
			val = creator("")
		} else {
			val = creator("")
		}
		value = reflect.ValueOf(val)
		fmt.Println(value)
		fmt.Println(value.Pointer())
	}
}

func TestPtrs9(t *testing.T) {
	var vals []func() string
	var value reflect.Value

	vals = getters("0")
	value = reflect.ValueOf(vals[0])
	fmt.Println(value)
	fmt.Println(value.Pointer())

	vals = getters("1")
	value = reflect.ValueOf(vals[0])
	fmt.Println(value)
	fmt.Println(value.Pointer())
}

func TestPtrs10(t *testing.T) {
	f1 := creator("f1")
	f2 := creator("f2")
	fmt.Printf("%s; %T; %x\n", f1(), f1, reflect.ValueOf(f1).Pointer())
	fmt.Printf("%s; %T; %x\n", f2(), f2, reflect.ValueOf(f2).Pointer())

	fmt.Println(reflect.DeepEqual(f1, f2))
}

func TestPtrs11(t *testing.T) {
	f0 := creator("run-0")
	f1 := creator("run-1")
	fmt.Println(reflect.ValueOf(f0))
	fmt.Println(reflect.ValueOf(f1))
	fmt.Println(reflect.ValueOf(f0).Pointer())
	fmt.Println(reflect.ValueOf(f1).Pointer())
}

func TestPtrs12(t *testing.T) {
	f0 := getters("run-0")[0]
	f1 := getters("run-1")[0]
	fmt.Println(reflect.ValueOf(f0))
	fmt.Println(reflect.ValueOf(f1))
	fmt.Println(reflect.ValueOf(f0).Pointer())
	fmt.Println(reflect.ValueOf(f1).Pointer())
}
