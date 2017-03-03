package main

import "fmt"

//var x = 0

func main() {
	x := 0
	increment := func() int { //anonymous func
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

}
