package main

import (
	"fmt"
)

type persona struct {
	nombre, apellido string
	heladosFavoritos []string
}

func main() {
	p1 := persona{
		nombre:           "Tomás",
		apellido:         "Cabello",
		heladosFavoritos: []string{"Oreo", "KitKat", "Vainilla"},
	}

	p2 := persona{
		nombre:           "Jaime",
		apellido:         "Garrido",
		heladosFavoritos: []string{"Fresa", "Dulce de Leche", "Turrón"},
	}

	mapa := map[string]persona{
		"Cabello": p1,
		"Garrido": p2,
	}

	if v, ok := mapa["Cabello"]; ok {
		fmt.Println(v)
	}

	if v, ok := mapa["Garrido"]; ok {
		fmt.Println(v)
	}

	for k, v := range mapa {
		fmt.Println(k)
		fmt.Printf("\t%s %s\n", v.nombre, v.apellido)
		for i, helado := range v.heladosFavoritos {
			fmt.Printf("\t\t%d: %s\n", i, helado)
		}
	}
}
