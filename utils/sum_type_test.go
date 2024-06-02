package utils_test

import (
	"github.com/andrewwphillips/test_snip/utils"
	"testing"
)

func TestSumType(t *testing.T) {
	c := utils.NewCircle(2)
	d := c
	d.SetRadius(3.14)

	println(c == d)
}
