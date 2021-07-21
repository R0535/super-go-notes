package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id       int
	vacation bool
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.id = 1
	ftEmployee.age = 23
	ftEmployee.name = "Jose"
	ftEmployee.vacation = false
	fmt.Println(ftEmployee)
}
