package main

import "fmt"

func main() {
	a := 9
	b := 7
	c := 2
	d := 4
	e := 8
	f := 6
	g := 3
	h := 1
	fmt.Println((a > b) && (c < d))
	fmt.Println((e == e) || (f != f))
	fmt.Println(!(g <= h))
}
