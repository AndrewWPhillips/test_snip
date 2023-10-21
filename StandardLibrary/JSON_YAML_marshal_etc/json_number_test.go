package __

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// This is from https://www.gojek.io/blog/relooking-at-golangs-reflect-deepequal
// There was a long blog about DeepEqual etc
// Simple solution is just to declare Id as an int (not interface{})

type School struct {
	//Id   interface{}
	Id   int
	Name string
}

func TestName(t *testing.T) {
	x := School{
		Id:   1,
		Name: "Golang Public School",
	}
	bytes, err := json.Marshal(x)
	if err != nil {
		println("error while unmarshalling the json ", err)
		os.Exit(1)
	}
	var y School
	err = json.Unmarshal(bytes, &y)
	if err != nil {
		println("error while unmarshalling the json ", err)
		os.Exit(1)
	}
	if reflect.DeepEqual(x, y) {
		println("both x & y are same")
	} else {
		println("x & y are different")
	}
	fmt.Println("value of x is ", x)
	fmt.Println("value of y is ", y)
}
