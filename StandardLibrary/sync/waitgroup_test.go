package __

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		log.Print("Hello,")
		wg.Done()
	}()

	wg.Wait()
	log.Println("world")
}

func TestWGAddWrong(t *testing.T) {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		time.Sleep(time.Second)
		log.Print("Hello,")
		wg.Done()
	}()

	wg.Wait()
	log.Println("world")
}

func TestWGNegative(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	wg.Done()
	wg.Done() // panic
	wg.Wait()
	log.Println("world")
}

func TestWGPositive(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2)
	wg.Done() // counter is still 1
	wg.Wait() // wait forever or deadlock
	log.Println("world")
}
