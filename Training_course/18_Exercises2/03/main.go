package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Println((a && b) || (b && a) || !(b && b)) //true
}
