package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("No hay raíz cuadrada real de un número negativo: %v.", f)
		return 0, ErrNorgateMath
	}
	return 42, nil
}
