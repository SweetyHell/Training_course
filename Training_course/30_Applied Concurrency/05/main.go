package main

import (
	"fmt"
)

func main() {
	var mything []int
	for i := 1; i <= 10; i++ {
		mything = append(mything, i)

		for n := range fact(gen(mything...)) {
			fmt.Println(n)
		}
	}
}

func gen(nums ...int) chan float64 {
	out := make(chan float64)
	go func() {
		for _, n := range nums {
			out <- float64(n)
		}
		close(out)
	}()
	return out
}
func fact(in chan float64) chan float64 {
	out := make(chan float64)
	var fact float64
	fact = 1
	go func() {
		for n := range in {
			fact = fact * n

		}
		out <- fact
		close(out)
	}()
	return out
}
