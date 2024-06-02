package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("./sleep.pprof")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer pprof.StopCPUProfile()

	loop()
}

func loop() {
	for i := 0; i < 3; i++ {
		fmt.Println("Loop:", i+1)
		sleep10()
		sleep20()
		sleep30()
	}
}

func sleep10() {
	println("10")
	time.Sleep(10 * time.Millisecond)
}

func sleep20() {
	println("20")
	time.Sleep(20 * time.Millisecond)
}

func sleep30() {
	println("30")
	time.Sleep(30 * time.Millisecond)
}
