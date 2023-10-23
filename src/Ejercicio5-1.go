package main

import (
	"fmt"
)

type persona struct {
	nombre, apellido string
	heladosFavoritos []string
}

func main() {
	p := persona {
		nombre: "Tomás",
		apellido: "Cabello",
		heladosFavoritos: []string{"Oreo", "KitKat", "Vainilla"},
	}

	p2 := persona {
		nombre: "Jaime",
		apellido: "Garrido",
		heladosFavoritos: []string{"Fresa", "Dulce de Leche", "Turrón"},
	}

	fmt.Println(p)
	fmt.Println(p2)
	fmt.Println()

	fmt.Printf("Para %s %s, los mejores helados son:\n", p.nombre, 
		p.apellido)
	for i, v := range p.heladosFavoritos {
		fmt.Printf("\t%d\t%s\n", i, v)
	}

	fmt.Printf("Para %s %s, los mejores helados son:\n", p2.nombre, 
		p2.apellido)
	for i, v := range p2.heladosFavoritos {
		fmt.Printf("\t%d\t%s\n", i, v)
	}
}