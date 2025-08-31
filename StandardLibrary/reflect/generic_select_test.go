//go:build go1.22

package __

import (
	"reflect"
	"testing"
)

func GetFromChannels[T any](channels []<-chan T) (r T, ok bool) {
	cases := make([]reflect.SelectCase, 0, len(channels))
	for _, c := range channels {
		cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(c)})
	}
	if _, recv, recvOK := reflect.Select(cases); recvOK {
		return recv.Interface().(T), true
	}
	return
}

func GetFromChannelsNoBlock[T any](channels []<-chan T) (r T, ok bool) {
	cases := make([]reflect.SelectCase, 0, len(channels))
	for _, c := range channels {
		cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(c)})
	}

	// Add a "default" case so the select won't block if no channel is ready
	cases = append(cases, reflect.SelectCase{Dir: reflect.SelectDefault})
	if _, recv, recvOK := reflect.Select(cases); recvOK {
		return recv.Interface().(T), true
	}
	return
}

func TestGetFromChannels(t *testing.T) {
	channels := make([]<-chan int, 0, 8)
	for i := range 8 {
		ch := make(chan int, 3)
		ch <- i * 10
		ch <- i*10 + 1
		ch <- i*10 + 2
		channels = append(channels, ch)
	}

	for range 25 {
		v, ok := GetFromChannelsNoBlock(channels)
		println(v, ok)
		if !ok {
			break
		}
	}
}
