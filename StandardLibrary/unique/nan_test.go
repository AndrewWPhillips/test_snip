package __

import (
	"fmt"
	"math"
	"testing"
	"unique"
)

func TestUniqueNan(t *testing.T) {
	h1 := unique.Make(math.NaN())
	h2 := unique.Make(math.NaN())
	fmt.Println(h1, h2, h1 == h2)
}
