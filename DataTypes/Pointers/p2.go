// Golang program to demonstrate the declaration
// and initialization of pointers
package main

import "fmt"

func main() {

	var x int = 5748

	var p *int // pointer dec

	p = &x // pointer init

	fmt.Println("Value in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
}
