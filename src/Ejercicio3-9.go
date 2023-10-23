package main

import (
	"fmt"
)

func main() {
	var deporteFav string = "Baloncesto"
	switch deporteFav {
	case "Baloncesto":
		fmt.Println("Tu deporte favorito es el Baloncesto")
	case "Fútbol":
		fmt.Println("Tu deporte favorito es el Fútbol")
	case "Tenis":
		fmt.Println("Tu deporte favorito es el Tenis")
	default:
		fmt.Println("No tienes deporte favorito")
	}

}