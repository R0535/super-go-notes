package main

import (
	"fmt"
	"strconv"
	"time"
)

// func main() {
// 	var x int //declaracion dejandole saber que tipo es
// 	x = 8
// 	y := 7 //Declaracion inferida por el tipo de dato que llega
// 	fmt.Println(x)
// 	fmt.Println(y)

// 	//manejo de errores
// 	num, err := strconv.ParseInt("651561651", 0, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(num)
// 	}
// 	m := make(map[string]int) //[Llave]Valor
// 	m["Key"] = 6
// 	fmt.Println(m["Key"])

// 	s := []int{1, 2, 2, 3}
// 	for index, value := range s {
// 		fmt.Println(index, value)
// 	}

// 	s = append(s, 12)
// 	for index, value := range s {
// 		fmt.Println(index, value)
// 	}
// 	c:=make(chan int)//se crea un canal para comunicar rutinas	[1]
// 	go doSomething(c)//go es una palabra reservada que le dice que cree la rutina [2]
// 	<- c //Aqui regresa la senal enviada desde la rutina doSomething [7]
// 	g:=25
// 	fmt.Println(g)
// 	h := &g//apuntador con & es el espacio de memoria donde esta guardado g [output = 0x020xx026x0]
// 	fmt.Println(h)
// 	fmt.Println(*h)//sabiendo la referencia a la cajita que contiene g, con * desempaquetamos el contenido [output = 25]
// }

func doSomething(c chan int){//a esta funcion que se convierte en rutina por [go] le cae un canal de tipo entero  [3]
	time.Sleep(3*time.Second)//espera 3 segundos [4]
	fmt.Println("Done")//imprime que esta hecho		[5]
	c <- 1 // Al canal que nos trajimos, le devuelve una senal con el valor de 1		[6]
}
