package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy of s:", c)
}
