package main

import (
	"fmt"
	"sort"
)

type People []string

func (p People) Len() int           { return len(p) }
func (p People) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p People) Less(i, j int) bool { return p[i] < p[j] }
func main() {
	studygroup := People{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studygroup)
	sort.Sort(studygroup)
	fmt.Println(studygroup)
}
