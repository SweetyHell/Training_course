package main

import (
	"fmt"
)

func main() {
	n := Hashbucket("Go", 12)
	fmt.Println(n)
}
func Hashbucket(word string, buckets int) int {
	//letter:=rune(word[0]) - need to chenge type to int32
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
