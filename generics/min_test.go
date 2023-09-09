package __

import (
	"log"
	"testing"

	"golang.org/x/exp/constraints"
)

func MinInt(first int, rest ...int) (retval int) {
	retval = first
	for _, v := range rest {
		if v < retval {
			retval = v
		}
	}
	return
}

func TestMinInt(t *testing.T) {
	log.Println(MinInt(77, 24, 42, 99, 73, 101, 24))
	log.Println(MinInt(-1, 24, -2))
}

func Min[T constraints.Ordered](first T, rest ...T) (retval T) {
	retval = first
	for _, v := range rest {
		if v < retval {
			retval = v
		}
	}
	return
}

func TestMinGeneric(t *testing.T) {
	log.Println(Min(77, 24, 42, 99, 73, 101, 24))
	log.Println(Min(-1, 24, -2))
}
