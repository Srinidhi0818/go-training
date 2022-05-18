package main

import (
    "strconv"
    "fmt"
)

func main() {
	a := 123
	fmt.Printf("\nThe type of a is : %T", a)
    t := strconv.Itoa(a)
    fmt.Println(t)
	fmt.Printf("\nThe type of t is : %T", t)
}