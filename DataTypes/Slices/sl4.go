package main

import (
	"fmt"
)

func main() {
	var intSlice = make([]int, 10)        
	var strSlice = make([]string, 10, 20) 

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

}
