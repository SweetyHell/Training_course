package main

import "fmt"

const (
	A = iota
	B
	C
)
const (
	D = iota
	E
	F
)
const (
	_ = iota
	G = iota * 10
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
}
