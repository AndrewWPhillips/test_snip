package __

import (
	"fmt"
	"testing"
	"time"
)

type E struct {
	*time.Time
}

func TestEmbedPointer(t *testing.T) {
	now := time.Now()
	e := E{Time: &now}
	fmt.Println(e.Add(time.Hour))
}
