package __

import (
	"fmt"
	"testing"
)

type fooError string

func (e *fooError) Error() string { return string(*e) }

func what() *fooError { return nil }

func TestReturn(t *testing.T) {
	var err error = what()
	fmt.Println(err == nil)
}
