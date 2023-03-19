package __

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

// See https://golangbot.com/june-2019-quiz-results/
// I got 1 and 2 wrong because I did not read closely
// I got 11 wrong because I forgot when defer func args are evaluated

func hello() []string {
	return nil
}

func TestQuiz1(t *testing.T) {
	h := hello // h is func not slice
	if h == nil {
		log.Println("nil")
	} else {
		log.Println("not nil")
	}
}

func TestQuiz2(t *testing.T) {
	i := 2
	s := "1000"
	if len(s) > 1 {
		i, _ := strconv.Atoi(s)
		i = i + 5
	}
	fmt.Println(i) // 2
}

func hello2(num ...int) {
	num[0] = 18
}

func TestQuiz3(t *testing.T) {
	i := []int{5, 6, 7}
	hello2(i...)
	fmt.Println(i[0]) // 18
}

func hello3(i int) {
	fmt.Println(i)
}

func TestQuiz11(t *testing.T) {
	i := 5
	defer hello3(i) // 5 - args evaluated here
	i = i + 10
}

func TestQuiz13A(t *testing.T) {
	var m map[string]int
	delete(m, "h")
	fmt.Println(m["h"])
}
