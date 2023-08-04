package __

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorsIs(t *testing.T) {
	err := fmt.Errorf("This is the base error")
	fmt.Printf("type %T %q\n", err, err)
	err2 := fmt.Errorf("This is the derived error:%w", err)
	fmt.Printf("type %T %q\n", err2, err2)
	err3 := fmt.Errorf("doubly derived error:%w", err2)
	fmt.Printf("type %T %q\n", err3, err3)

	fmt.Println(errors.Is(err3, err))
	fmt.Println(errors.Is(err3, err2))
	fmt.Println(errors.Is(err3, err3))
	fmt.Println(errors.Is(err3, nil))

	var err4 any
	errors.As(err3, &err4)
	fmt.Printf("type %T %q\n", err4, err4)
}
