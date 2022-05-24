package main

import (
	"fmt"
)

type Persons struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var p Persons
	fmt.Println(p)

	p1 := Persons{"Rajeev", "Singh", 26}
	fmt.Println("Persons1: ", p1)
	p2 := Persons{"John", "Snow", 45}
	fmt.Println("Persons2: ", p2)
	p3 := Persons{FirstName: "Robert"}
	fmt.Println("Persons3: ", p3)
}
