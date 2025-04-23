package __

import (
	"fmt"
	"testing"
	"time"
)

func TestSliceAppendBug(t *testing.T) {
	var sTime, eTime []int64 // start and end times
	for range 2 {
		now := time.Now().Unix()
		sTime = append(sTime, now)
		eTime = append(sTime, now+10) // 10 secs later
	}
	fmt.Println(sTime[0] < eTime[0], sTime[1] < eTime[1])
}
