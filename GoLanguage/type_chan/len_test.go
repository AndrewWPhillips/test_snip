package __

import (
	"testing"
)

func TestLenChan(t *testing.T) {
	var ch chan int
	println(len(ch), cap(ch))
	ch = make(chan int, 5)
	println(len(ch), cap(ch))

	println(len([]int{}), len([]int(nil))) // 0 0
}
