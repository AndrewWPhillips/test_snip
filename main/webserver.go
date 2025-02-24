package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe("localhost:6060", nil)
	for i := 0; ; i++ {
		fmt.Println(i)
		ws_flame()
		//ws_graph()
		time.Sleep(3 * time.Second)
	}
}

func ws_flame() {
	r := 0
	for i := range [100000]struct{}{} {
		r += i
	}
}

func ws_graph() {
	for i := 0; i < 1e3; i++ {
		time.Sleep(time.Microsecond)
	}
}
