package main

import (
	"fmt"
)

func main() {
	var anoNac int = 1998
	var anosVividos int = 0

	fmt.Println("Has vivido los años:")
	fmt.Println(anoNac)

	for {
		if anoNac+anosVividos >= 2023 {
			break
		} 

		anosVividos++
		fmt.Println(anoNac + anosVividos)
	}

	fmt.Println()
	fmt.Println("Has vivido", anosVividos, "años")
}
