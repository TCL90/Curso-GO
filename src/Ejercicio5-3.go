package main

import (
	"fmt"
)

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func main() {
	c1 := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "Azul",
		},
		cuatroRuedas: true,
	}

	s1 := sedan{
		vehiculo: vehiculo{
			puertas: 5,
			color:   "Rojo",
		},
		lujoso: true,
	}

	fmt.Println(c1)
	fmt.Println(s1)

	fmt.Printf("El camión tiene %d puertas, es de color %s y %t\n",
		c1.puertas, c1.color, c1.cuatroRuedas)
	fmt.Printf("El sedán tiene %d puertas, es de color %s y %t",
		s1.puertas, s1.color, s1.lujoso)
}
