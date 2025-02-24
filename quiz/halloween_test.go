package __

import (
	"fmt"
	"testing"
	"time"
)

func TestQ1(t *testing.T) {
	var treasure map[string]string
	treasure["island"] = "ead man's chest"
	treasure["tomb"] = "mummie's gold"
	//treasure := map[string]string{
	//	"island": "dead man's chest",
	//	"tomb":   "mummie's gold",
	//}

	fmt.Println(treasure)
}

func TestQ1ClearMap(t *testing.T) {
	treasure := map[string]string{}
	clear(treasure)
	fmt.Println(treasure == nil)
}

func TestQ3(t *testing.T) {
	creatures := map[int]string{
		0: "Vampires",
		1: "Werewolves",
		2: "Zombies",
	}
	for _, it := range creatures {
		fmt.Println("You are being chased by", it)
		break
	}
}
func unlucky(num ...int) {
	num[0] = 13
}

func TestQ4(t *testing.T) {
	i := []int{7, 8, 9}
	unlucky(i...)
	fmt.Println(i[0]) // 13
}

func day(yymmdd string) (r time.Time) {
	r, _ = time.Parse("060102", yymmdd)
	return
}

func TestQTime(t *testing.T) {
	events := map[time.Time]string{
		day("241031"): "Halloween",
		day("241225"): "Christmas",
	}
	check := time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC)
	for when, name := range events {
		if check == when {
			println("It's", name)
		}
	}
}

func TestTime2(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Round(0))
	y, m, d := now.Date()
	fmt.Println(y, m, d)
	fmt.Println(y, int(m), d)
}

func CheckDate(toCheck time.Time) {
	events := map[string]string{"241031": "Halloween", "241225": "Christmas"}
	for yymmdd, name := range events {
		date, _ := time.Parse("060102", yymmdd)
		if toCheck.Equal(date) {
			fmt.Println("It's", name)
		}
	}
}

func TestTime3(t *testing.T) {
	//CheckDate(time.Date(2024, 10, 31, 0, 0, 0, 0, time.Local))
	//CheckDate(time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC))
	CheckDate(time.Date(2024, 10, 31, 0, 0, 0, 0, time.Local).UTC())
}

func TestTime4(t *testing.T) {
	now := time.Now() // local time
	utc := now.UTC()
	local := utc.Local()
	println(now == utc, now == local)
	println(now.Equal(utc), now.Equal(local))
}

func TestWitchingHour(t *testing.T) {
	//midnight := time.Date(2024, 10, 31, 0, 0, 0, 0, time.Local)
	midnight := time.Now()
	witchingHour := midnight.Round(1)
	fmt.Println(midnight == witchingHour)
	fmt.Println(midnight.Equal(witchingHour))
}

func ethereal() error {
	return nil
}

func TestQ5(t *testing.T) {
	if f := ethereal; f == nil {
		fmt.Println("No error")
	} else {
		fmt.Println("Something mysterious")
	}
}

var zombie = "alive"

func isAlive(frankenstein string) {
	fmt.Println("Frank is", frankenstein, "zombie is", zombie)
}

func TestQ6(t *testing.T) {
	var frankenstein = "alive"
	defer isAlive(frankenstein)
	frankenstein, zombie = "dead", "dead"
}

func TestNilSlice(t *testing.T) {
	var s []int
	copy(s, []int{1, 2, 3})
	s[0] = 13
	println(s)
}

func TestQ2(t *testing.T) {
	//var spooks []string
	spooks := make([]string, 3)
	copy(spooks, []string{"ghost", "fairy", "invisible man"})
	fmt.Println(spooks)
}

func TestClear(t *testing.T) {
	var s []int
	clear(s)
	fmt.Println(s, s == nil)
}

func TestSliceRealloas(t *testing.T) {
	dead := []string{"ghost", "zombie"}
	alive := dead
	alive = append(alive, "a", "b", "c", "d")
	alive[0] = "z"
	fmt.Println(dead[0])
}

func TestSliceResize(t *testing.T) {
	living := []string{"zombie", "Lazarus"}
	dead := living[0:0] // make them dead
	dead = dead[:5]
	println(dead[0])
	// runtime error: index out of range [0] with length 0
}

func TestConstNum(t *testing.T) {
	const insane = 1.3
	rational := 1.3
	fmt.Println(insane*insane == rational*rational)
}

func TestConstNum2(t *testing.T) {
	const insane = 1e1_000_000
	fmt.Printf("%.20f", insane/1e1_000_000)
}

type Ghost struct {
	a int
}

type Ghola struct {
	a int `json:"a"`
}

func TestGhola(t *testing.T) {
	g1 := Ghost{a: 13}
	g2 := Ghola(g1)
	fmt.Println(g2.a)
}
