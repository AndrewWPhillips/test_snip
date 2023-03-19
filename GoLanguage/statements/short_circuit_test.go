package __

import (
	"log"
	"testing"
)

func TestShortCircuit(t *testing.T) {
	var f func() bool
	//f = func() bool {
	//	return false
	//}
	if f != nil && f() { // if f is nil then calling f() is not attempted since (false && X == false)
		log.Println("Yes")
	} else {
		log.Println("No")
	}
}
