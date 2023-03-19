package __

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNulByteInString(t *testing.T) {
	orig := []string{"AA", "BB"}
	sep := string('\x00')
	got := strings.Split(strings.Join(orig, sep), sep)
	fmt.Println(len(got), got)
	if !reflect.DeepEqual(got, orig) {
		t.FailNow()
	}
}
