package __

import (
	"fmt"
	"testing"
)

func f98() (r *[]byte) {
	{
		var password []byte
		//defer clear(password)  // WRONG
		password = []byte("secret") // get password from somewhere
		defer clear(password)
		r = &password
	}
	return
}

func TestNested2(t *testing.T) {
	fmt.Println(*f98())
}
