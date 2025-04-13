package pkg1

import _ "unsafe"

var (
	A = 42
	b = 43
	c = 44
)

func Pub() int { return 41 }

// priv1 is a private functions used to demonstrate "take" access
func priv1() int {
	return b
}

// Give access to priv2 to (any) "main" package
//
//go:linkname priv2 main.priv2
func priv2() int {
	return c
}
