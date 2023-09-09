package main

func main() {
	i := 42
	escaper(&i)
}

var q *int

func escaper(p *int) {
	q = p
}
