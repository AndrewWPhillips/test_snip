package main

import (
	"fmt"
)

func main() {
	for i := 10; i < 20; i++ {
		fmt.Printf("%p\n", &i)
	}
}
