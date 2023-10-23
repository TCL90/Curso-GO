package main

import (
	"fmt"
)

type persona struct {
	nombre, apellido string
	edad             int
}

func (p persona) presentar() {
	fmt.Println("Hola soy", p.nombre, p.apellido, 
		"y mi edad es", p.edad, "años.")
}

func main() {
	p := persona{
		nombre:   "Tomás",
		apellido: "Cabello",
		edad:     25,
	}

	p.presentar()
}
