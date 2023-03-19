package __

import (
	"log"
	"testing"
)

func TestShadowing(t *testing.T) {
	one := -1
	log.Printf("one is at %p\n", &one)
	one, two := 1, 2 // can use := with one because two is new
	log.Printf("one now at %p\n", &one)
	{
		one, two := 3, 4 // can use := because they are both new (in nested scope)
		log.Println(one, two)
	}
	log.Println(one, two)
}
