package __

import (
	"errors"
	"log"
	"testing"
)

func errgen() (int, error) {
	return 1, errors.New("new error")
}

func TestIgnoreError(t *testing.T) {
	i, _ := errgen() // ignoring errors is almost as easy as C
	log.Println(i)
}
