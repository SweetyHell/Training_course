package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type Doublezero struct {
	Person
	Licence bool
}

func (p Person) fullname() string {
	return p.First + " " + p.Last
}
func (p Doublezero) fullname() string {
	return "hello, " + p.First + " " + p.Last
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   20,
	}
	p2 := Doublezero{
		Person: Person{
			First: "Orlando",
			Last:  "Palando",
			Age:   34,
		},
		Licence: false,
	}

	fmt.Println(p1)
	fmt.Println(p1.First)
	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
