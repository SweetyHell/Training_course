package main

import (
	"fmt"
	"sort"
)

func main() {
	type People []string
	studygroup := People{"Zeno", "John", "Al", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Jenny", "Joseph"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Strings(studygroup)
	fmt.Println("sorted:", studygroup)
	sort.StringSlice(s).Sort() //or sort.Sort(sort.Stringslice(s)) //we can use this because Stringslice is implementing an interface that is needed for Sort
	fmt.Println("sorted:", s)
	sort.Ints(n)
	fmt.Println("sorted:", n)
	fmt.Println("---------------")
	sort.Sort(sort.Reverse(sort.StringSlice(studygroup)))
	fmt.Println("sorted:", studygroup)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("sorted:", s)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("sorted:", n)
}
