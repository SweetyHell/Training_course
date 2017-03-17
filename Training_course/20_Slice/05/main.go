package main

import (
	"fmt"
)

func main() {
	myslice := make([]int, 1)
	fmt.Println(myslice[0])
	myslice[0] = 7
	myslice = append(myslice, 100) //adds next element 100
	fmt.Println(myslice)
	myslice[0]++ //adds1 to the value
	myslice[1]++ //+= 1
	fmt.Println(myslice[0], myslice[1])
}
