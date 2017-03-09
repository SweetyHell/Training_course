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
	delete(n, 3)
	if val, exists := n[3]; exists {
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("The value doesn't exist")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}
	fmt.Println("map:", n)
}
