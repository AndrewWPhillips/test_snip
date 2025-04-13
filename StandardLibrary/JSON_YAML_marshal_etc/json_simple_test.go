package __

import (
	"encoding/json"
	"log"
	"testing"
	"unique"
)

func TestSimpleJSON(t *testing.T) {
	type customer struct {
		Name string
		//Vouchers []string
		Vouchers []unique.Handle[string]
	}

	h1 := unique.Make("abc")
	h2 := unique.Make("def")

	got, err := json.Marshal(
		customer{
			Name: "John Doe",
			//Vouchers: []string{"abc", "def"},
			Vouchers: []unique.Handle[string]{h1, h2},
		})
	if err != nil {
		log.Fatalln("json.Marshall error:", err)
	}
	log.Printf("got JSON %q\n", got)
}
