package __

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringsSplit(t *testing.T) {
	fmt.Println(strings.Split("a,b,c", ",")[0])
}

func TestStringsSplitAfter(t *testing.T) {
	fmt.Println(strings.SplitAfter("a,b,c", ",")[0])
}

func TestSplitSeq(t *testing.T) {
	for _, s := range strings.Split("abc,;def,z", ",;") {
		println(s)
	}
}
