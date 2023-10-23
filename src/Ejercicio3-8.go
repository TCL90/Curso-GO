package main

import (
	"fmt"
)

func main() {
	switch {
	case 2 == 2:
		fmt.Println("2 es igual a 2")
	case 3 == 4:
		fmt.Println("Esto no se imprime")
	default:
		fmt.Println("Por defecto")
	}	
}