package main

import "fmt"

func addNumbers(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func main() {

	result := addNumbers(1, 3)

	fmt.Println("Sum:", result)
}
