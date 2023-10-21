package __

import (
	"log"
	"testing"
)

// TestDeferReturn can change named return value
func TestDeferReturn(t *testing.T) {
	f := func() (ret int) {
		defer func() { ret = 42 }() // ret = 42
		return 1                    // ret = 1
	}
	log.Println(f()) // 42
}

// TestDeferReturnWrong can NOT change return value if not named
func TestDeferReturnWrong(t *testing.T) {
	f := func() int { // URC (unnamed return value) is int
		ret := 1
		defer func() {
			ret = 42 // ret = 42 but URV still = 1
		}()
		return ret // URV = 1
	}
	log.Println(f()) // 1
}
