package __

import (
	"log"
	"testing"
)

func TestFallthrough(t *testing.T) {
	i := 42
	b := true
	switch i {
	case 1:
		log.Println("1")

	case 42:
		log.Println("42")
		//fallthrough //can't put fallthrough here
		if !b {
			//fallthrough // can't put fallthrough here
		}
		fallthrough // can put fallthrough here

	default:
		log.Println("default")

	}
}
