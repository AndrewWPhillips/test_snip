package __

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBitShift(t *testing.T) {
	s := "4ADDF6C259EBAFF8"
	if i, err := strconv.ParseUint(s, 16, 64); err == nil {
		i := i>>25 + 0x1008000400
		fmt.Printf("%x\n", i)
	}
}
