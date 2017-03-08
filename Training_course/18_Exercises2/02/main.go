package main

import "fmt"

func main() {
	//var values = []float64{}
	a := maximum(10, 1, 211, 111, 1222)
	//	fmt.Scan(values...)

	fmt.Println(a)
	//maximum(values...)
}

func maximum(ss ...int) int {

	m := 0
	for _, v := range ss {
		if m < v {
			m = v
		}
	}
	return m
}
