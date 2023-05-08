package __

import (
	"errors"
	"log"
	"testing"
)

type MyError struct{}

func (MyError) Error() string { return "my error" }

func TestNilOrNot(t *testing.T) {
	var myErr *MyError = nil
	var err error
	err = myErr
	log.Println(err == nil)
}

func stub() error { return errors.New("TOE") }
func TestQuizFunc(t *testing.T) {
	result := stub
	if result != nil {
		log.Println(result)
		return
	}
	log.Println("OK")
}
