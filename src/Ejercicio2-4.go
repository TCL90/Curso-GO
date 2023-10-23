package main

import (
	"fmt"
)

func main() {
	var a int = 42
	fmt.Printf("El número %d es:\n-%d en decimal.\n-%b en binario.\n-%#x en hexadecimal.\n\n",
		a, a, a, a)
		
	fmt.Println("Hacemos Bit Shifting una posición a la izquierda a", a)
	a = a << 1

	// Hacer este shift es como multiplicar por 2 elevado a 1. 42 x 2 = 84
	fmt.Printf("El número %d es:\n-%d en decimal.\n-%b en binario.\n-%#x en hexadecimal.\n",
		a, a, a, a)
}
