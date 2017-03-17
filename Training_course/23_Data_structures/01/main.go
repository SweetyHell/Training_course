package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}
type doublezero struct {
	person
	First         string
	Licencetokill bool
}

func main() {
	p1 := doublezero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "double zero seven",
		Licencetokill: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.person.First)
	fmt.Println(p1.First)
}
