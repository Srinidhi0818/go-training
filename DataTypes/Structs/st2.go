package main

import (
	"fmt"
)

type Idol struct {
	Name    string
	Group   string
	Country string
}

func main() {
	c := Idol{
		Name:    "Jin",
		Group:   "BTS",
		Country: "Kor",
	}
	fmt.Println("Idol Name: ", c.Name)
	fmt.Println("Idop Country: ", c.Country)

	c.Country = "WWH"
	fmt.Println("Idol: ", c)
}
