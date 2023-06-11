package __

import (
	"expvar"
	"log"
	"strconv"
	"testing"
)

type MyInt int

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

func TestExpVar(t *testing.T) {
	var mi MyInt = 42
	expvar.Publish("x", mi)
	log.Println(expvar.Get("x"))
}
