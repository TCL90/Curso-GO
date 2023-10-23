package main

import (
	"fmt"
	"strings"
)

func main() {
	mapa := map[string][]string{
		"Tomás_Cabello":  []string{"Bici", "Baloncesto", "Juegos de mesa"},
		"Jaime_Garrido":  []string{"Atletismo", "Natación", "Senderismo"},
		"Felipe_Limones": []string{"Danza", "Fútbol", "Beisbol"},
	}

	fmt.Println(mapa)
	fmt.Println(mapa["Tomás_Cabello"])
	fmt.Println(mapa["Jaime_Garrido"])
	fmt.Println(mapa["Felipe_Limones"])
	fmt.Println()

	for k, v := range mapa {
		k = strings.Replace(k, "_", " ", 1)
		fmt.Println("Usuario:", k)
		fmt.Printf("\t Hobbies:\n")

		for _, hobbie := range v {
			fmt.Printf("\t\t%s\n", hobbie)
		}

		fmt.Println()
	}
}
