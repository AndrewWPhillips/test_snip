package __

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"
)

// See https://golangbot.com/june-2019-quiz-results/
// I got 1 and 2 wrong because I did not read closely
// I got 11 wrong because I forgot when defer func args are evaluated

func stub() error {
	return errors.New("ERROR")
}

func TestQuiz1(t *testing.T) {
	result := stub
	if result != nil {
		log.Println(result())
		return
	}
	log.Println("OK")
}

func TestQuiz2(t *testing.T) {
	i, s := 1, "100"
	if len(s) > 1 {
		i, _ := strconv.Atoi(s)
		i++
	}
	fmt.Println(i) // 1
}

func ten(num ...int) {
	num[0] = 10
}

func TestQuiz3(t *testing.T) {
	i := []int{5, 6, 7}
	ten(i...)
	fmt.Println(i[0]) // 10
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
