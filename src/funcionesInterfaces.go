package main

import (
	"fmt"
)

type persona struct {
	nombre, apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido, "el agente secreto se presenta")
}

func (p persona) presentarse() {
	fmt.Println("Hola, soy", p.nombre, p.apellido, "la persona se presenta")
}

type humano interface {
	presentarse()
}

func bar(h humano) {
	fmt.Println("Fui pasado a la función bar", h)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Tomás",
			apellido: "Cabello",
		},
		lpm: true,
	}

	p := persona{
		nombre:   "Jaime",
		apellido: "Garrido",
	}

	fmt.Println(as1)
	as1.presentarse()

	bar(p)
	bar(as1)
}
