package main

import "fmt"

type Person struct {
	First string
	Age   int
}

// func (p Person) fullname() string {
// 	return p.First + " " + p.Last
// }

func main() {
	p1 := &Person{"James", 20}

	fmt.Println(*p1)
	fmt.Println(p1)
	fmt.Println(p1.First)
	fmt.Printf("%T\n", p1)
	//fmt.Println(p1.fullname())

}
