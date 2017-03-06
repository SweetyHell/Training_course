package main

import "fmt"

func main() {
	var sn, bn int

	fmt.Println("Enter big number:")
	fmt.Scan(&bn)
	fmt.Println("Enter small number:")
	fmt.Scan(&sn)
	fmt.Println("Remainder is: ", bn%sn)
}
