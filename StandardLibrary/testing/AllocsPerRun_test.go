package testing

import (
	"log"
	"testing"
)

func TestAllocsPerRun(t *testing.T) {
	a := testing.AllocsPerRun(10, func() {
		for i := 0; i < 10; i++ {
			_ = make(chan int) // 1 alloc here
			//log.Println(&struct{ int }{1}) // 2 allocs?
		}
	})

	log.Println(a)

}
