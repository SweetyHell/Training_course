package main

import "fmt"

func main() {

	n := map[int]string{
		1: "Alena",
		2: "Olena",
		3: "Hellena",
		4: "Helen",
	}
	n[5] = "Hell"

	for key, val := range n {
		fmt.Println("key ", key, "value ", val)
	}
}
