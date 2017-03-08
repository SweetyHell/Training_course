package main

import "fmt"

func main() {
	var x int
	fmt.Print("Enter a number:")
	fmt.Scan(&x)
	myfunction1(x)
}

func myfunction1(y int) {
	var z bool
	if y%2 == 0 {
		z = true
	}
	fmt.Println(y/2, z)

}

/*
Todd's version:
func half(n int) (int, bool){
return n/2,n%2==0

func main(){
h,even:=half(5)
fmt.Println(h,even)
}
}
*/
