package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	p1 := persona{
		nombre:   "Tomás",
		apellido: "Cabello",
		edad:     25,
	}

	p2 := persona{
		nombre:   "Jaime",
		apellido: "Garrido",
		edad:     42,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.nombre, p1.apellido, p1.edad)
	fmt.Println(p2.nombre, p2.apellido, p2.edad)

}
