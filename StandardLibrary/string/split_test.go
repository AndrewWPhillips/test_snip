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

func TestStringsSplitSeq(t *testing.T) {
	for s := range strings.SplitSeq("a,b,c", ",") {
		fmt.Println(s)
	}
}
