package main

import (
	"fmt"
	"github.com/andrewwphillips/test_snip/private/pkg1"

	//_ "github.com/andrewwphillips/test_snip/private/pkg1"
	_ "unsafe"
)

// Take access of pkg1 private 'priv1' function as pkg1Priv1
//
//go:linkname pkg1Priv1 github.com/andrewwphillips/test_snip/private/pkg1.priv1
func pkg1Priv1() int // this definition must match declaration in pkg1 !!
//func pkg1Priv1() int8  // works (on little-endian systems only?)

//go:linkname priv2
func priv2() int

func main() {
	fmt.Println(pkg1.Pub())
	fmt.Println(pkg1.A)
	fmt.Println(pkg1Priv1())
	fmt.Println(priv2())
}
