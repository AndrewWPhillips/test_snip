package __

import (
	"log"
	"testing"
)

func Min(first int, rest ...int) (retval int) {
	retval = first
	for _, v := range rest {
		if v < retval {
			retval = v
		}
	}
	return
}

func TestMinInt(t *testing.T) {
	log.Println(Min(77, 24, 42, 99, 73, 101, 24))
	log.Println(Min(-1, 24, -2))
}
