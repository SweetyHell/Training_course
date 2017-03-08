package main

import "fmt"

func main() {

	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(ss ...int) {
	fmt.Print(ss)
	fmt.Printf("%T \n", ss)
}
