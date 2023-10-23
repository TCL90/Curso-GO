package main

import (
	"fmt"
)

func main() {
	b := []string{"Batman", "Jefe", "Vestido de negro"}
	r := []string{"Robin", "Ayudante", "Vestido de colores"}

	arrayMulti := [][]string{b, r}

	for i, reg := range arrayMulti {
		fmt.Println(i)
		for j, val := range reg {
			fmt.Printf("\tÍndice: %d\t%s\n", j, val)
		}
	}
}
