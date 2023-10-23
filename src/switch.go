package main

import (
	"fmt"
)

func main() {
	switch {
	case 2 == 4, 4 ==8, 9 == 9:
		fmt.Println("No debería imprimirse, ahora sí")
	case 3 == 3:
		fmt.Println("Debería imprimirse")
		fallthrough
	case 4 == 5:
		fmt.Println("No debería imprimirse2")
	default:
		fmt.Println("Imprimiendo default")
	}

	switch "Manzana" {
	case "Pera", "Limón":
		fmt.Println("Frutas verdes")
	case "Manzana", "Ciruela", "Fresa":
		fmt.Println("Frutas rojas")
		fallthrough
	default:
		fmt.Println("Frutas")
	}
}