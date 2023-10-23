package main

import "fmt"

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

type humano interface {
	hablar()
}

func (p *persona) hablar() {
	fmt.Println("Hablo como una persona.")
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		Nombre:   "Tomás",
		Apellido: "Cabello",
		Edad:     25,
	}

	/*
		p2 := persona{
			Nombre:   "Jaime",
			Apellido: "Garrido",
			Edad:     34,
		}
	*/

	diAlgo(&p1)
	//diAlgo(p2)
}
