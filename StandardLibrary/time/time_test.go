package time

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTimeUnix(t *testing.T) {
	tt := time.Now()
	log.Println(tt.String())
	tu := time.Now().Unix()
	tt = time.Unix(tu, 1)
	log.Println(tt.String(), tu)
}

func Stringify(s []time.Time) (retval []string) {
	for _, v := range s {
		retval = append(retval, v.String())
	}
	return
}

//type Stringer struct {}
//func (s Stringer) String() string { return "" }
//
//func Stringify(type Stringer)(s []Stringer) (ret []string) {
//	for _, v := range s {
//		ret = append(ret, v.String())
//	}
//	return ret
//}

func TestDateStringer(t *testing.T) {
	vals := []time.Time{time.Now(), time.Now().Add(25 * time.Hour)}
	strs := Stringify(vals)
	log.Println(strs)
}

// TestDateAdjusted shows that passing a day of zero to time.Date gives the previous day
func TestDateAdjusted(t *testing.T) {
	tm := time.Date(2000, 1, 0, 0, 0, 0, 0, time.Local)
	fmt.Println(tm) // 1999-12-31 (day before 2001-01-01)
}

func TestNowUtcLocal(t *testing.T) {
	log.Println(time.Now())
	log.Println(time.Now().Local())
	log.Println(time.Now().UTC())

	// underlying time is not affected by timezone
	log.Println(time.Now().Unix())
	log.Println(time.Now().Local().Unix())
	log.Println(time.Now().UTC().Unix())
}
