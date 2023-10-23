package main

import (
	"fmt"
)

type persona struct {
	nombre, apellido string
	edad             int
}

func cambiame(p *persona) {
	(*p).nombre = "Martín"
	(*p).apellido = "Kafw"
	p.edad = 42
}

func main() {
	p := persona{
		nombre:   "Tomás",
		apellido: "Cabello",
		edad:     25,
	}

	fmt.Println("Persona antes de cambiame", p)
	cambiame(&p)
	fmt.Println("Persona después de cambiame", p)
}
