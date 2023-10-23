package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	x := sum("Hola", xi...)
	fmt.Println("El total almacenado es", x)
}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(s)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el índice", i, "suma", v, "al total, quedando", suma)
	}

	fmt.Println("El total es", suma)
	return suma
}