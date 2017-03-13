package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name    string
	Surname string //`json:"-"`
	Age     int    //`json:"changed_Age"`
	notexp  int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"Name":"James", "Surname":"Bond","Age": 22}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)

}
