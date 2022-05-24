package main

import (
	"fmt"
)

func main() {

	var x int64 = 55

	var y float64 = float64(x)

	fmt.Printf("x = %d \n", x)
	fmt.Printf("y = %f \n", y)

	fmt.Printf("decimal 3 = %.3f", y)

}
