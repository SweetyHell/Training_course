package main

import "fmt"

func main() {
	for i := 100; i < 110; i++ {
		fmt.Println([]byte(string("alena")), `
     `, rune(i))
		//	fmt.Printf("%T \n", i)
		//	fmt.Printf("%T \n", rune(i))
	}

}
