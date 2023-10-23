package main

import (
	"fmt"
)

type persona struct {
	nombre, apellido string

}

type agenteSecreto struct {
	persona,
	lpm bool
}

func (a agenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido)
}

func main() {
	as1 := agenteSecreto {
		persona: persona{
			nombre: "Tomás"
			apellido: "Cabello"
		}
		lpm: true
	}

	fmt.Println(as1)
	as1.presentarse()
}

