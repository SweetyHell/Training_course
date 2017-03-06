package main

import "fmt"

func visit(numbers []int, callback1 func(int)) {
	for _, n := range numbers {
		callback1(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// callback: passing a func as an argument
