package main

import "fmt"

type Employee struct {
	id int
}
type Person struct {
	name string
	age  int
}

type FullTimeEmployee struct {
	Person
	Employee
	EndDate int
}
type PartTimeEmployee struct {
	Person
	Employee
	taxRate int
}

//Interface
type PrintInfo interface {
	getMessage() string
}

//Reciever functions 
//para implementar interfaces, las receivers no llevan {*}
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "FTE"
}
func (ptEmployee PartTimeEmployee) getMessage() string {
	return "PTE"
}

//implementations of interface over structs
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	fte := FullTimeEmployee{}
	pte := PartTimeEmployee{}
	getMessage(fte)
	getMessage(pte)
}
