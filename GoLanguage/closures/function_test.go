package __

import (
	"log"
	"testing"
)

type User struct{ Name string }

func NoModify(u *User) { u = &User{Name: "Paul"} }
func Modify(u *User)   { u.Name = "Jessica" }
func ByValue(u User)   { u.Name = "Eric" }

func TestPassPointer(t *testing.T) {
	u := User{Name: "Leto"}
	log.Println(u.Name)
	NoModify(&u)
	log.Println(u.Name)
	Modify(&u)
	log.Println(u.Name)
	ByValue(u)
	log.Println(u.Name)
}

func TestNilFunc(t *testing.T) {
	var f func(int)
	// f(1) // panics
	log.Printf("%p\n", f) // 0x0 (why not <nil>?)
}

func fyyy(i int) { log.Printf("yyy %d\n", i) }

func TestCallback(t *testing.T) {
	var f func(int)

	if f == nil {
		log.Printf("Nil function pointer\n")
	}
	f = fyyy
	f(42)
}
