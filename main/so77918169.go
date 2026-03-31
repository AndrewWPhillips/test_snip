package main

import (
	"log"
	"net"
	"os"
	"runtime/debug"
	"runtime/trace"
	"time"
)

var quit = make(chan struct{})

func printGcStats() {
	s := debug.GCStats{}
	s.PauseQuantiles = make([]time.Duration, 101)

	lastGc := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	maxPause := time.Duration(0)

	for range time.Tick(time.Second) {
		debug.ReadGCStats(&s)
		if s.LastGC.Equal(lastGc) || len(s.Pause) == 0 {
			continue
		}
		for i, t := range s.PauseEnd {
			if t.Compare(lastGc) > 0 {
				if s.Pause[i] > maxPause {
					maxPause = s.Pause[i]
				}
				if s.Pause[i] > time.Millisecond*50 {
					log.Printf("GC %d  maxPause: %v", s.NumGC, maxPause)
					quit <- struct{}{}
					return
				}
			} else {
				break
			}
		}
		log.Printf("GC %d  maxPause: %v", s.NumGC, maxPause)
		lastGc = s.LastGC
	}
}

func init() {
	go printGcStats()
}

func write() {
	conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{Port: 9104})
	buf := make([]byte, 512)
	for {
		conn.Write(buf)
	}
}

var rChan = make(chan []byte, 1000)

func read() {
	for x := range rChan {
		_ = x
	}
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	ln, _ := net.ListenUDP("udp", &net.UDPAddr{Port: 9104})
	go write()
	go read()
	go func() {
		for {
			buf := make([]byte, 512)
			ln.ReadFrom(buf)
			rChan <- buf
		}
	}()

	<-quit
}
