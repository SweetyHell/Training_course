package main

import "fmt"

const (
	_  = iota
	kB = 1 << (iota * 10)
	mB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%b\t", kB)
	fmt.Printf("%d\n", kB)
	fmt.Printf("%b\t", mB)
	fmt.Printf("%d\n", mB)
}
