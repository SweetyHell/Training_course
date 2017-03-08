package main

import "fmt"

func main() {
	records := make([][]string, 0)
	stud1 := make([]string, 4)
	stud1[0] = "Olena"
	stud1[1] = "Nov"
	stud1[2] = "100"
	stud1[3] = "99"
	records = append(records, stud1)
	stud2 := make([]string, 4)
	stud2[0] = "Alex"
	stud2[1] = "Kerg"
	stud2[2] = "90"
	stud2[3] = "95"
	records = append(records, stud2)
	fmt.Println(records)
}
