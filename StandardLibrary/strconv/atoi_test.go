package __

import (
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestAtoi(t *testing.T) {
	log.Println(strconv.Atoi("0"))                          // 0
	log.Println(strconv.Atoi("1"))                          // 1
	log.Println(strconv.Atoi(" 0"))                         // 0, 0 strconv.Atoi: parsing "0": invalid syntax
	log.Println(strconv.Atoi(" 1"))                         // 0, 0 strconv.Atoi: parsing "0": invalid syntax
	log.Println(strconv.Atoi(strings.TrimSpace(" 1 ")))     // 1
	log.Println(strconv.Atoi(strings.TrimLeft(" 1", " ")))  // 1
	log.Println(strconv.Atoi(strings.TrimLeft(" 1a", " "))) // 0, 0 strconv.Atoi: parsing "1a": invalid syntax
	log.Println(strconv.Atoi(strings.Trim(" 1a", " a")))    // 1
	log.Println(strconv.Atoi("a"))                          // 0, 0 strconv.Atoi: parsing "a": invalid syntax
	log.Println(strconv.Atoi(""))                           // 0 strconv.Atoi: parsing "": invalid syntax
}
