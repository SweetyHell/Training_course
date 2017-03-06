package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello,W!")
	}
	greeting()
	fmt.Printf("%T", greeting)
}
