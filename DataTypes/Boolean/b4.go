package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a bool

	s := "True"
	a, _ = strconv.ParseBool(s)

	fmt.Println(a)
}
