package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo()())
}

func foo() func() int{
	return func() int{
		fmt.Println("Hola, soy la función que retorna foo.")
		return 2
	}
}