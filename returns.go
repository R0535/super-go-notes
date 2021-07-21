package main

import (
	"fmt"
)

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}

}

func getValues(x int) (double int, triple int, quadruple int) {
	double = x * 2
	triple = x * 3
	quadruple = x * 4
	return
}
func main() {
	fmt.Println(sum(12, 13, 4, 5, 6))
	printNames("Jose", "Eduardo")
	fmt.Println(getValues(1))
}
