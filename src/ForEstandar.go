package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	//Ciclo anidado
	for i := 0; i < 10; i++ {
		fmt.Printf("Entramos en el ciclo externo %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tEntramos en el ciclo interno %d\n", j)
		}
	}
}
