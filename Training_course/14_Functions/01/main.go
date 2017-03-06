package main

import "fmt"

func main() {
	greet("Jane")
	greet2("John", "von Hoff")
}

func greet(name string) {
	fmt.Println(name)
}

func greet2(name string, surname string) {
	fmt.Println(name, surname)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
