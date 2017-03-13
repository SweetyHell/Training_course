package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Surname string `json:"-"`
	Age     int    `json:"changed_Age"`
	notexp  int
}

func main() {
	p1 := Person{"James", "Bond", 22, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Printf("%T\n", bs)
}
