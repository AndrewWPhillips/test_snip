package __

import (
	"log"
	"testing"
)

type readableThrift interface{}

type factory func(string) readableThrift

func fact(s string) readableThrift {
	switch s {
	case "int":
		p := new(int)
		*p = 42
		return p
	}
	return nil
}

func TestFactory(t *testing.T) {
	var f factory
	f = fact
	val := f("int")

	v := val.(*int)
	log.Println(*v)
}
