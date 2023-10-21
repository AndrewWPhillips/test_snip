package main

func main() {
	//for i := 0; i < 10; i++ {
	//	go println(i) // race detector does not detect this (something to do with println?)
	//}
	j := 0
	for i := 0; i < 10; i++ {
		j += i
	}
}
