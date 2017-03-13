package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	notexp int
}

func main() {
	p1 := Person{"James", 22, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
	fmt.Printf("%T\n", bs)
}
