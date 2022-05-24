package main

import "fmt"

func main() {
	var name string = "srinidhi"
	var pointer *string = &name

	fmt.Println("name :", name)
	fmt.Println("pointer :", pointer)

	fmt.Println("*pointer :", *pointer)
}
