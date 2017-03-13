package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Surname string //`json:"-"`
	Age     int    `json:"changed_Age"`
	//notexp  int
}

func main() {
	var p1 Person
	fmt.Println(p1)
	//p1 := Person{"James", "Bond", 22, 007}

	bs := []byte(`{"Name":"James","Surname":"Bond","changed_Age":20}`)
	json.Unmarshal(bs, &p1)
	fmt.Println("-------------------------")

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)

}
