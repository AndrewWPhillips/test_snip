package ___test

import (
	"fmt"
	"testing"

	__ "github.com/andrewwphillips/test_snip/GoLanguage/type_struct"
)

func TestName(t *testing.T) {
	p := __.NewPerson("andrew", 67)
	p2 := *p
	fmt.Println(p2.Name)
	p2.Name = ""
	fmt.Println(p2.Name)
}
