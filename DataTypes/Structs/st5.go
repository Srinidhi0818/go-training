package main

import "fmt"

type Point struct {
	X float64
	Y float64
}

func main() {
	p1 := Point{10, 20}
	p2 := p1
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)

	p2.X = 15
	fmt.Println("\nModified p2:")
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)
}
