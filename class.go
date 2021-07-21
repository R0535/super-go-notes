package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}
func (e *Employee) SetName(name string) {
	e.name = name
}
func (e *Employee) GetId() int {
	return e.id
}
func (e *Employee) GetName() string {
	return e.name
}
func main() {
	e := Employee{}
	e.id = 10
	e.name = "Juan"
	fmt.Println(e)
	e.SetId(11)
	fmt.Println(e)
	e.SetName("Pepe")
	fmt.Println(e)
}
