package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 5
	// y := func() int {
	// 	return x * 2
	// }()
	// fmt.Println(y)
	c := make(chan int)//Con go routines
	go func() {
		fmt.Println("Auto destruccion automatica iniciada")
		time.Sleep(5 * time.Second)
		c <- 1
	}()
	<-c
}
