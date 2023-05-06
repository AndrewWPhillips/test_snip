package __

import (
	"log"
	"testing"
)

func TestForLoopVarDeclaredOutside(t *testing.T) {
	var i int
	for i = 0; i < 4; i++ {
		log.Println(i)
	}
	log.Println(i)
}

func TestRangeVarDeclaredOutside(t *testing.T) {
	var i int
	list := []string{"ZERO", "ONE", "TWO", "THREE"}
	for i = range list {
		log.Println(i)
	}
	log.Println(i)
}
