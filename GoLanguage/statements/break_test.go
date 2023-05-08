package __

import (
	"log"
	"testing"
)

func TestBreakInSwitch(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			if i > 1 {
				break // goes to end of switch not for loop
			}
			log.Println("two")
		default:
			log.Println(i)
		}
	}
}
