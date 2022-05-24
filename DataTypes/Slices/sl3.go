package main

import "fmt"

func main() {

	slices := []string{"Hi", "All", "Good", "Day"}
	fmt.Println(len(slices))
	fmt.Println(cap(slices))
	fmt.Println(slices)
}
