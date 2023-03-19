package __

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestValue(t *testing.T) {
	type t1 struct{}
	type t2 struct{}

	ca := context.Background()
	cb := context.WithValue(ca, t1{}, "B")
	cc := context.WithValue(cb, t2{}, "C")

	log.Println("ca", ca.Value(t1{}), ca.Value(t2{}))
	log.Println("cb", cb.Value(t1{}), cb.Value(t2{}))
	log.Println("cc", cc.Value(t1{}), cc.Value(t2{}))
}

// TextCheckCancel tests checking for a cancelled context in a loop
func TestCheckCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(20 * time.Nanosecond)
		cancel()
	}()
	//log.Println(ctx.Deadline())
	for i := 0; i < 1e9; i++ {
		if err := ctx.Err(); err != nil {
			log.Println(err)
			break
		}
		if i%1000 == 0 {
			log.Println(i)
		}
	}
}

// TextCheckDone tests checking for a cancelled context using the Done channel
func TestCheckDone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Nanosecond)

	ch := make(chan int, 1)
	i := 1
loop:
	for {
		select {
		case v := <-ch:
			log.Println("Got:", v)
		case ch <- i:
			log.Println("Wrote:", i)
			i++
		case <-ctx.Done():
			log.Println("Context cancelled")
			break loop
		}
	}
	cancel() // just to be safe
}
