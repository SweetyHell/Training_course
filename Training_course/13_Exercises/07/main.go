package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {

		switch {
		case i%3 == 0, i%5 == 0:
			sum = sum + i
			// case i%3 == 0:
			// 	fmt.Println("Fizz")
			// case i%5 == 0:
			// 	fmt.Println("Buzz")
			// default:
			// 	fmt.Println(i)
			//
		}

	}
	fmt.Println(sum)
}
