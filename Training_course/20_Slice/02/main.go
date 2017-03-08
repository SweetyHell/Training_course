package main

import "fmt"

func main() {

	myslice2 := make([]int, 0, 5)

	//	fmt.Printf("%T", myslice)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	for i := 0; i < 50; i++ {
		myslice2 = append(myslice2, i)
		fmt.Println("Len:", len(myslice2), "Capacity:", cap(myslice2), "Value:", myslice2[i])
	}

}
