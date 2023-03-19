package __

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSliceSelect(t *testing.T) {
	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	go func() {
		cases := []reflect.SelectCase{
			{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch1)},
			{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch2)},
		}
		for {
			chosen, value, OK := reflect.Select(cases)
			fmt.Println(chosen, value, OK)
			if !OK {
				cases[chosen].Chan = reflect.ValueOf(nil)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		if i == 5 {
			close(ch1)
			ch1 = nil // signal it's closed
		} else if ch1 != nil {
			ch1 <- i
		}
		ch2 <- i * 10
	}
}
