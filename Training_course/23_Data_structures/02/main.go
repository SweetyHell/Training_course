package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doublezero struct {
	person
	Licence bool
}

func (p person) fullname() string {
	return p.First + " " + p.Last
}
func (p doublezero) fullname() string {
	return "hello, " + p.First + " " + p.Last
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   20,
	}
	p2 := doublezero{
		person: person{
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
