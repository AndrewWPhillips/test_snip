//go:build go1.22
// +build go1.22

package main

// Uses CPU profiling to show string builder is faster than just using += on a string

import (
	"fmt"
	"github.com/pkg/profile"
	"strconv"
	"strings"
)

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()
	var str string
	for range 1_000_000 {
		str = bs()
	}
	fmt.Println(str)
}

// Takes about 10,000 nsecs
//func bs() string {
//	str := ""
//	for i := range 100 {
//		str += strconv.Itoa(i)
//	}
//	return str
//}

// Takes about 700 nsecs
func bs() string {
	var sb strings.Builder
	sb.Grow(2 * 100)
	for i := range 100 {
		sb.WriteString(strconv.Itoa(i))
	}
	return sb.String()
}
