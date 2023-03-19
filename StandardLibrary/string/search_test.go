package __

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

const s = "Andrew Phillips"

func TestIndexFunc(t *testing.T) {
	log.Println(strings.IndexFunc(s, func(r rune) bool { return r == 'A' })) // 0
	log.Println(strings.IndexFunc(s, func(r rune) bool { return r == 'e' })) // 4
	log.Println(strings.IndexFunc(s, func(r rune) bool { return r == 'z' })) // -1

	fmt.Println(strings.Count(s, "i")) // 2
	fmt.Println(strings.Count(s, ""))  // 16 unicode code points
}

func TestContains(t *testing.T) {
	fmt.Println(strings.Contains(s, "drew"))      // true
	fmt.Println(strings.Contains(s, "and"))       // false
	fmt.Println(strings.ContainsAny(s, "uvwxyz")) // true
	fmt.Println(strings.ContainsAny(s, "xyz"))    // false

	fmt.Println(strings.Index(s, "ips"))    // 12
	fmt.Println(strings.IndexAny(s, "ips")) // 9
	fmt.Println(strings.IndexByte(s, 'i'))  // 9
}

func TestFields(t *testing.T) {
	fmt.Printf("%q\n", strings.Fields(s))                                            // ["Andrew" "Phillips"]
	fmt.Printf("%q\n", strings.FieldsFunc(s, func(r rune) bool { return r == 'i' })) // ["Andrew Ph" "ll" "ps"]
}

func TestSplit(t *testing.T) {
	x := strings.Split(s, "i")
	fmt.Println(strings.Join(x, " ")) // Andrew Ph ll ps
}
