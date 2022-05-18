package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "HI"
	c := "hello"
	b := strings.ToLower(a)
	fmt.Println("Lowered a:", b)
	d := strings.ToUpper(c)
	fmt.Println("Uppered c:", d)
}
