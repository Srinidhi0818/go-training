package main

import (
	"fmt"
	"strconv"
)

func main() {
	var val1 bool = true
	var val2 bool = false
	var str string
	str = strconv.FormatBool(val1)
	fmt.Println(str)

	str = strconv.FormatBool(val2)
	fmt.Println(str)
}
