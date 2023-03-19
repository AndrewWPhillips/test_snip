package __

import (
	"log"
	"reflect"
	"testing"
)

func TestTypeName(t *testing.T) {
	i := 0
	log.Println(reflect.TypeOf(i))
	log.Println(reflect.TypeOf(2))
	log.Println(reflect.TypeOf(&i))
	log.Println(reflect.TypeOf(struct{ a int }{3}))
	log.Println(reflect.TypeOf(TestTypeName))
}

// TestInterface test getting the type of an interface
func TestInterface(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)
	var i interface{}
	i = ch1
	log.Println(reflect.TypeOf(i)) // chan int
	i = ch2
	log.Println(reflect.TypeOf(i)) // chan string

	// An interface is implicitly cast to an interface{} (aka any) containing the value in the original interface
	// To get the "real" type of variable i, we get the value of a ptr to i, then the type of what it points to
	v := reflect.TypeOf(&i)
	log.Println(v.Elem()) // interface{}
}
