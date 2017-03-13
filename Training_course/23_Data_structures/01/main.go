package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}
type Doublezero struct {
	Person
	First         string
	Licencetokill bool
}

func main() {
	p1 := Doublezero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "double zero seven",
		Licencetokill: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.Person.First)
	fmt.Println(p1.First)
}
