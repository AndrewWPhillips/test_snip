package __

import (
	"testing"
)

type A int

func (a A) inc() A {
	return a + 1
}

func TestAlias(t *testing.T) {

}
