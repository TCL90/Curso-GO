package main

import (
	"fmt"
)

func main() {
	coche := struct {
		puertas   int
		color     string
		deportivo bool
		marca     string
		amigos    map[string]int
		canciones []string
	}{
		puertas:   5,
		color:     "Blanco",
		deportivo: true,
		marca:     "Mercedes",
		amigos: map[string]int{
			"Fran":    678980976,
			"Antonio": 658097488,
		},
		canciones: []string{
			"I belong with you",
			"We will never back together",
		},
	}

	fmt.Println(coche)
	fmt.Printf("El %s tiene %d puerta(s) y es de color %s\n", coche.marca,
		coche.puertas, coche.color)
	fmt.Println("Lo tienen tus amigos:")

	for k, v := range coche.amigos {
		fmt.Printf("\tNombre: %s, teléfono: %d\n", k, v)
	}

	fmt.Println("\nLa lista de reproducción incluye:")
	for i, v := range coche.canciones {
		fmt.Printf("\t%d: %s\n", i, v)
	}

}
