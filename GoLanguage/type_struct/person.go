package __

// struct used by info_hide_test.go
// Person is in package __ and so private fields are hidden from package ___test

type person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *person {
	return &person{
		Name: name,
		Age:  age,
	}
}
