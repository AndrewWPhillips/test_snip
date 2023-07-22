package __

// tests various functions that return a modified string

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	const s = "Andrew Phillips"
	lower := strings.ToLower(s)
	log.Println(lower)                                                                      // andrew phillips
	log.Println(strings.Title(lower))                                                       // Andrew Phillips
	log.Println(strings.ToTitle(lower))                                                     // ANDREW PHILLIPS
	log.Println(strings.ToUpper(lower))                                                     // ANDREW PHILLIPS
	log.Println(strings.Trim(lower, "adeilnps"))                                            // rew ph
	log.Println(strings.TrimFunc(lower, func(r rune) bool { return r == 'a' || r == 's' })) // ndrew phillip

	log.Println(strings.TrimPrefix(lower, "and")) // rew phillips
	log.Println(strings.TrimPrefix(lower, "adn")) // andrew phillips

}

func TestBuilder(t *testing.T) {
	var builder strings.Builder
	builder.Grow(128)
	fmt.Fprintf(&builder, "<%s> %g", "me", math.Pi) // *strings.Builder implements io.Writer
	log.Println(builder.Len(), builder.String())    // 22 <me> 3.141592653589793

	reader := strings.NewReader(builder.String())
	log.Println(reader.ReadRune()) // 60 1 <nil>
}

func TestReplacer(t *testing.T) {
	replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	log.Println(replacer.Replace("< < <"))
	log.Println(replacer.Replace("This is <b>HTML</b>!"))
}

func TestScanner(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(`one
two
three
four
`))
	var (
		text []byte
		n    int
	)
	for scanner.Scan() {
		n++
		text = append(text, fmt.Sprintf("%d. %s\n", n, scanner.Text())...)
	}
	//os.Stdout.Write(text) // What is the question?
	fmt.Println(string(text))
}
