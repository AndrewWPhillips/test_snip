package __

import (
	"errors"
	"fmt"
	"log"
	"sort"
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
	i, s := 13, "13"
	if len(s) > 1 {
		i, _ := strconv.Atoi(s)
		i++
	}
	fmt.Println(i) // 13
}

func ten(num ...int) {
	num[0] = 10
}

func TestQuiz3(t *testing.T) {
	i := []int{7, 8, 9}
	ten(i...)
	fmt.Println(i[0]) // 10
}

func hello3(i int) {
	fmt.Println(i)
}

func TestQuiz11(t *testing.T) {
	sort.Search()
	i := 13
	defer hello3(i) // 13 - args evaluated here
	i = i + 10
}

func TestQuiz13A(t *testing.T) {
	var m map[string]int
	delete(m, "h")
	fmt.Println(m["h"])
}

// TestConstFloat demonstrates how compile-time constants are evaluated exactly.  That is,  f2*f2 is
// evaluated at run-time, f1*f1 uses IEEE float64 (base-2 exponent) so cannot represent 1.3 exactly
func TestConstFloat(t *testing.T) {
	const f1 = 1.3
	var f2 = 1.3
	fmt.Println(f1*f1 == f2*f2)
}

// What does this print?
// a. true
// b. false
