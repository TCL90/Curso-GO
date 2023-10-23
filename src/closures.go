package main

import (
	"fmt"
)

var x int //scope nivel de paquete

func main() {
	fmt.Println(x)

	x++
	fmt.Println(x)

	foo()
	fmt.Println(x)

	{
		y := 42 // scope solo dentro de las llaves
		fmt.Println(y)
	}

	v := incrementor()
	b := incrementor()
	fmt.Println(v())
	fmt.Println(v())
	fmt.Println(v())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	fmt.Println("Hola")
	x++
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}