// Program to illustrate code reusability in function

package main

import "fmt"

func sumSquare(num int) {
	square := num * num
	fmt.Printf("Square of %d is %d\n", num, square)
}

func main() {

	sumSquare(3)
	sumSquare(5)
	sumSquare(10)

}
