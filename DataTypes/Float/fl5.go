package main

import (
	"fmt"
)

func main() {
	floval := 11.21
	s := fmt.Sprintf("%f", floval)
	fmt.Printf("%T, %v\n", s, s)
}
