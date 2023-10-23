package main

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("Hola")
	fmt.Println("Que")
	fmt.Println("Tal?")
}

func foo() {
	defer func() {
		fmt.Println("Foo diferida se ejecutó.")
	}()
	fmt.Println("Soy la función Foo.")
}