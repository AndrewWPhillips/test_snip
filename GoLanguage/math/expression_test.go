package __

import (
	"go/types"
	"log"
	"testing"
)

// TestExpr is to test evaluating expressions in the debugger
func TestExpr(t *testing.T) {
	a := 1
	u := types.Universe.Lookup("print")
	log.Println(a, u)
}
