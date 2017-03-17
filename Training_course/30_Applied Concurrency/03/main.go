package main

import (
	"fmt"
)

func main() {

	c := factorial(4)
	//fmt.Println(<-c) // this also works instead of the below range!
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(a int) chan int {
	out := make(chan int)
	fact := 1
	go func() {
		for i := 1; i <= a; i++ {
			fact = fact * i
		}
		out <- fact
		close(out)

	}()
	return out
}
