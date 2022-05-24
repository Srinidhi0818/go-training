package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("empty string:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("values in s:", s)
	fmt.Println("fetching the value of particular s value:", s[2])

	fmt.Println("length of s:", len(s))

}
