package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

func main() {
	pEmp := new(Employee)

	pEmp.Id = 2926
	pEmp.Name = "Srinidhi"

	fmt.Println(pEmp)
}
