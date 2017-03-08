package main

import "fmt"

func main() {
	myslice := []string{"a", "b", "c", "d", "e"}
	myslice2 := make([]int, 50, 100)
	fmt.Println(myslice)
	//	fmt.Printf("%T", myslice)
	fmt.Println(myslice[2:4])
	fmt.Println(myslice[:])
	fmt.Println(myslice2[:5])
}
