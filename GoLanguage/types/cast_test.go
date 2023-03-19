package __

import (
	"log"
	"testing"
)

func TestIntToBool(t *testing.T) {
	var b bool
	c := 1
	//b = bool(c)
	b = c != 0
	log.Println(b)
}
