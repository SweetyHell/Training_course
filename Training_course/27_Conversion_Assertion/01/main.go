package main

import "fmt"

func main() {
	var x = 12
	var y = 12.12222
	fmt.Println(y + float64(x))                     //converting x
	fmt.Println(int(y) + x)                         //converting y
	fmt.Println(string([]byte{'h', 'y', 'l', 'e'})) //converting []bytes to string ([]bytes containing runes)
	fmt.Println([]byte("Hello"))                    //convert string to slice of bytes
}
