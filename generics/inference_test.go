package __

import (
	"sync"
	"testing"
)

func Max2[T int | float64](values []T) T {
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func TestTwoTypes(t *testing.T) {
	//Max2(nil)    // can't infer if []int or []float64
	Max2[int](nil) // panics
}

func Max1[T int](values []T) T {
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
	sync.Mutex
}

func TestOneType(t *testing.T) {
	Max1(nil) // inferred to be []int(nil) since the type parameter can only be int
}
