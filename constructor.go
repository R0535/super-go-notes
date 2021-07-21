package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee { //Aca arriba se indica que va a regresar el Valor del contenedor de un empleado
	return &Employee{ //Aqui se crea un nuevo contenedor de empleado con sus atributos
		id:       id,
		name:     name,
		vacation: vacation,
	} //Se devuelve
}
func main() {
	//1
	e1 := Employee{}
	fmt.Println(e1)
	//2
	e2 := Employee{
		id:       1,
		name:     "Juan",
		vacation: true,
	}
	fmt.Println(e2)
	//3
	e3 := new(Employee) //este devuelve el apuntador
	e3.id = 2
	e3.name = "Pancho"
	e3.vacation = true
	fmt.Println(*e3)
	//4
	e4 := NewEmployee(3, "Pedro", false)
	fmt.Println(e4)
}
