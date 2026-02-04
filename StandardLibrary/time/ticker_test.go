package __

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	<-time.NewTimer(time.Second).C
}
