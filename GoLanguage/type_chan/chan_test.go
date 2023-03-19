package __

import (
	"fmt"
	"log"
	"testing"
)

// TestChain chains together channels to generate prime numbers
// Generates prime numbers by creating a "filter" for each new prime found & adding
// it to the end of then chain. If a number makes it through then it is a prime!
func TestChain(t *testing.T) {
	const MAX = 10000

	// Create generator for ints starting at the first prime (2)
	gen := make(chan int) // channel that "generates" the integers
	go func() {
		for i := 2; i < MAX; i++ {
			gen <- i
		}
		close(gen)
	}()

	var end <-chan int = gen // current end of chain of filters
	for {
		// TODO: handle close of start of chain
		prime := <-end     // get next prime
		fmt.Println(prime) // display the prime

		// Make a filter for this prime - go-routine with in/out channels
		next := make(chan int)
		go func(value int, src <-chan int, dst chan<- int) {
			for i := range src {
				if i%value > 0 {
					dst <- i
				}
			}
		}(prime, end, next)
		end = next
	}
}

// TestCompare checks that the same channel compares equal even for input/output channels
func TestCompare(t *testing.T) {
	c1 := make(chan int)
	c2 := c1
	c3 := chan<- int(c1)
	c4 := (<-chan int)(c1)
	d := make(chan int)
	log.Println(c1 == c2, c1 == c3, c1 == c4, c1 == d)
}

func TestIdentify(t *testing.T) {
	c1 := make(chan int)
	c2 := c1
	c3 := chan<- int(c1)
	d := make(chan int)

	log.Printf("%p %p %p %p\n", c1, c2, c3, d) // all have name "address" except d
}
