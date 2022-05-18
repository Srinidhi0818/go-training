package main

import (
	"fmt"
)

func main() {
	var a, b = 4, 5
	var r1 = (a + b) * (a + b) / 2

	a++
	fmt.Println("Increment of A:", a)

	b += 10
	fmt.Println("Increment of B:", b)
	fmt.Printf("res1 : %v", r1)
}
