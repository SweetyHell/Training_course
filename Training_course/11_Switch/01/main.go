package main

import "fmt"

func main() {

	switch "Medhi" {
	case "A":
		println("ok A")
	case "B", "D":
		println("ok B")
	case "Medhi":
		println("ok Medhi")
		fallthrough
	case "C":
		println("ok C")
	default:
		fmt.Println("no")

	}
}
