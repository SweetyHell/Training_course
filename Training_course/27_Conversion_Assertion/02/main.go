package main

import "fmt"

func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string) //interface type needs assortion, not convertion.
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
	fmt.Printf("%T\n", name)
	name = 12
	str2, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str2)
	} else {
		fmt.Printf("value is not a string\n")
	}
	fmt.Printf("%T\n", name)
}
