package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("La variable x es del tipo %T y valor %v\n", x, x)
	fmt.Printf("La variable y es del tipo %T y valor %v\n", y, y)
	fmt.Printf("La variable z es del tipo %T y valor %v\n", z, z)

	// Las valores asignados por defecto a las variables cuando
	// se declaran son los valores cero.
}